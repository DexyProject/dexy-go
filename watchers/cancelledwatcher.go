package watchers

import (
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/consumers"
)

type CancelledWatcher struct {
	orderbook orderbook.OrderBook

	in <-chan *consumers.CancelledMessage
}

func NewCancelledWatcher(book orderbook.OrderBook, in <-chan *consumers.CancelledMessage) CancelledWatcher {
	return CancelledWatcher{
		orderbook: book,
		in:        in,
	}
}


func (cw *CancelledWatcher) Watch() {
	for {
		msg := <- cw.in
		go cw.handle(msg)
	}
}

func (cw *CancelledWatcher) handle(msg consumers.Message) {
	hash := types.Hash{}
	hash.SetBytes(msg.(*consumers.CancelledMessage).Hash.Bytes[:])

	ok := cw.orderbook.RemoveOrder(hash)
	if !ok {
		msg.Reject()
		return
	}

	msg.Ack()
}