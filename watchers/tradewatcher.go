package watchers

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/consumers"
)

type TradeWatcher struct {
	history   history.History
	exchange  *exchange.ExchangeInterface
	orderbook orderbook.OrderBook

	in <-chan *consumers.TradedMessage
}

func NewTradeWatcher(history history.History, exchange *exchange.ExchangeInterface, book orderbook.OrderBook, in <- chan *consumers.TradedMessage) TradeWatcher {
	return TradeWatcher{
		history: history,
		exchange: exchange,
		orderbook: book,
		in: in,
	}
}

func (tf *TradeWatcher) Watch() {

	for {
		msg := <-tf.in
		tx := msg.Transaction

		err := tf.history.InsertTransaction(tx)
		if err != nil {
			msg.Reject()
			// @todo handle
			return
		}

		filled, err := tf.orderFilledAmount(tx.Maker, tx.OrderHash)
		if err != nil {
			msg.Reject()
			// @todo
			return
		}

		if tf.isOrderFilled(tx.OrderHash, filled) {
			tf.orderbook.RemoveOrder(tx.OrderHash) // @todo check response
			msg.Ack()
			return
		}

		tf.orderbook.UpdateOrderFilledAmount(tx.OrderHash, filled) // @todo check response
		msg.Ack()
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
