package watchers

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
)

type TradeWatcher struct {
	history   history.History
	exchange  exchange.ExchangeInterface
	orderbook orderbook.OrderBook
}

func (tf *TradeWatcher) Watch() error {

	sink := make(chan *exchange.ExchangeInterfaceTraded)

	_, err := tf.exchange.WatchTraded(nil, sink, make([][32]byte, 0))
	if err != nil {
		// @todo return
		return err // @todo better
	}

	for {

		trade := <-sink

		tx := types.Transaction{
			TransactionID: types.Bytes{Bytes: trade.Raw.TxHash.Bytes()},
			OrderHash: trade.Hash,
			BlockNumber: trade.Raw.BlockNumber,
			Timestamp: 0,
			Taker: types.Address{Address: trade.Taker},
			Maker: types.Address{Address: trade.Maker},
			Give: types.Trade{
				Token: types.Address{Address: trade.TokenGive},
				Amount: trade.AmountGive,
			},
			Get: types.Trade{
				Token: types.Address{Address: trade.TokenGet},
				Amount: trade.AmountGet,
			},
		}

	}

	return nil
}

// @todo this can probably use some optimization
func (tf *TradeWatcher) handleTransaction(transaction types.Transaction) {

	err := tf.history.InsertTransaction(transaction)
	if err != nil {
		// @todo handle
		return
	}

	filled, err := tf.orderFilledAmount(transaction.Maker, transaction.OrderHash)
	if err != nil {
		// @todo
		return
	}

	if tf.isOrderFilled(transaction.OrderHash, filled) {
		tf.orderbook.RemoveOrder(transaction.OrderHash) // @todo check response
		return
	}

	tf.orderbook.UpdateOrderFilledAmount(transaction.OrderHash, filled)
}

func (tf *TradeWatcher) isOrderFilled(order types.Hash, amount types.Int) (bool) {
	o := tf.orderbook.GetOrderByHash(order)
	return o.Get.Amount.Cmp(&amount.Int) == 0
}

func (tf *TradeWatcher) orderFilledAmount(maker types.Address, order types.Hash) (types.Int, error) {
	f, err := tf.exchange.Filled(nil, maker.Address, order)
	if err != nil {
		return types.Int{}, err
	}

	return types.Int{Int: *f}, nil
}
