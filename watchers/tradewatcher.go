package watchers

import (
	"log"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/subscribers"
	"github.com/DexyProject/dexy-go/types"
	"encoding/json"
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
		// @todo async
		msg, err := tf.subscriber.Listen()
		if err != nil {
			// @todo
			continue
		}

		var tx *types.Transaction
		err = json.Unmarshal([]byte(msg), tx)
		if err != nil {
			// @todo
			continue
		}

		go tf.handleTransaction(*tx)

		// @todo
		log.Print(msg)
	}

	// @todo read for trade events,
	return nil
}

func (tf *TradeWatcher) handleTransaction(transaction types.Transaction) {

	err := tf.history.InsertTransaction(transaction)
	if err != nil {
		// @todo handle
		return
	}

	tf.orderbook.UpdateOrderFilledAmount(transaction.OrderHash, transaction.Get.Amount)

	// @todo delete if amount == filled
}
