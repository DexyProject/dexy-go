package watchers

import (
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/types"
)

type TradeWatcher struct {
	history history.History
}

func (tf *TradeWatcher) Watch() error {
	// @todo read for trade events,
	return nil
}

func (tf *TradeWatcher) HandleTransaction(transaction types.Transaction) {

	err := tf.history.InsertTransaction(transaction)
	if err != nil {
		// @todo handle
		return
	}

	// @todo updated filled values
	// @todo delete if amount == filled

}
