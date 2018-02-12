package watchers

import (
	"encoding/json"
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
		// @todo async
		// @todo ack, dec message depending on if it worked.
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
		// @todo delete
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
