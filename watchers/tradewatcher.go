package watchers

import (
	"errors"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"go.uber.org/zap"
)

type TradeWatcher struct {
	history   history.History
	exchange  *contracts.Exchange
	orderbook orderbook.OrderBook

	in <-chan *consumers.TradedMessage
}

func NewTradeWatcher(history history.History, exchange *contracts.Exchange, book orderbook.OrderBook, in <-chan *consumers.TradedMessage) TradeWatcher {
	return TradeWatcher{
		history:   history,
		exchange:  exchange,
		orderbook: book,
		in:        in,
	}
}

func (tf *TradeWatcher) Watch() {
	for {
		msg := <-tf.in
		tx := msg.Transaction

		err := tf.history.InsertTransaction(tx)
		if err != nil {
			log.Error("tx rejected insert", zap.Error(err))
			msg.Reject()
			continue
		}

		filled, err := tf.orderFilledAmount(tx.OrderHash)
		if err != nil {
			log.Error("tx rejected filled amount", zap.Error(err))
			msg.Reject()
			continue
		}

		err = tf.handleFill(tx, filled)
		if err != nil {
			log.Error("tx rejected handle", zap.Error(err))
			msg.Reject()
			continue
		}

		msg.Ack()
	}
}

func (tf *TradeWatcher) handleFill(tx types.Transaction, filled types.Int) error {
	if !tf.isOrderFilled(tx.OrderHash, filled) {
		return tf.orderbook.UpdateOrderFilledAmount(tx.OrderHash, filled)
	}

	ok := tf.orderbook.RemoveOrder(tx.OrderHash)
	if !ok {
		return errors.New("failed to delete order")
	}

	return nil
}

func (tf *TradeWatcher) isOrderFilled(order types.Hash, amount types.Int) bool {
	o := tf.orderbook.GetOrderByHash(order)
	return o.Take.Amount.Cmp(&amount.Int) == 0
}

func (tf *TradeWatcher) orderFilledAmount(order types.Hash) (types.Int, error) {
	f, err := tf.exchange.Filled(nil, order.Hash)
	if err != nil {
		return types.Int{}, err
	}

	return types.Int{Int: *f}, nil
}
