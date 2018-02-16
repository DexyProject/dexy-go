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

	in <-chan types.Transaction
}

func NewTradeWatcher(history history.History, exchange exchange.ExchangeInterface, book orderbook.OrderBook, in <- chan types.Transaction) TradeWatcher {
	return TradeWatcher{
		history: history,
		exchange: exchange,
		orderbook: book,
		in: in,
	}
}

func (tf *TradeWatcher) Watch() {

	for {
		tx := <-tf.in

		// @todo we should be able to ack or reject, depending on if we fail somewhere

		err := tf.history.InsertTransaction(tx)
		if err != nil {
			// @todo handle
			return
		}

		filled, err := tf.orderFilledAmount(tx.Maker, tx.OrderHash)
		if err != nil {
			// @todo
			return
		}

		if tf.isOrderFilled(tx.OrderHash, filled) {
			tf.orderbook.RemoveOrder(tx.OrderHash) // @todo check response
			return
		}

		tf.orderbook.UpdateOrderFilledAmount(tx.OrderHash, filled)
	}
}

func (tf *TradeWatcher) isOrderFilled(order types.Hash, amount types.Int) bool {
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
