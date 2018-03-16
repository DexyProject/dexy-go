package watchers

import (
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/gateways/orderbook"
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
		msg := <-cw.in
		go cw.handle(msg)
	}
}

func (cw *CancelledWatcher) handle(msg consumers.Message) {
	ok := cw.orderbook.RemoveOrder(msg.(*consumers.CancelledMessage).Hash)
	if !ok {
		msg.Reject()
		return
	}

	msg.Ack()
}
