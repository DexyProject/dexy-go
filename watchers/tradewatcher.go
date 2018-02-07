package watchers

import (
	"log"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/subscribers"
	"github.com/DexyProject/dexy-go/types"
)

type TradeWatcher struct {
	subscriber subscribers.Subscriber
	history    history.History
	exchange   exchange.ExchangeInterface
	orderbook  orderbook.OrderBook
}

func (tf *TradeWatcher) Watch() error {

	err := tf.subscriber.Subscribe()
	if err != nil {
		return err
	}

	for {
		msg, err := tf.subscriber.Listen()
		if err != nil {
			// @todo
			continue
		}

		// @todo
		log.Print(msg)
	}

	// @todo read for trade events,
	return nil
}

func (tf *TradeWatcher) HandleTransaction(transaction types.Transaction) {

	err := tf.history.InsertTransaction(transaction)
	if err != nil {
		// @todo handle
		return
	}

	tf.orderbook.UpdateOrderFilledAmount(transaction.OrderHash, transaction.Get.Amount)

	// @todo delete if amount == filled
}
