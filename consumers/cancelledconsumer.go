package consumers

import (
	"log"
	"sync"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/event"
)

type CancelledConsumer struct {
	sync.Mutex

	exchange *exchange.ExchangeInterface

	out  chan<- *CancelledMessage
	stop chan struct{}

	ack    chan types.Bytes
	reject chan types.Bytes

	sub event.Subscription
}

func NewCancelledConsumer(ex *exchange.ExchangeInterface, out chan<- *CancelledMessage) CancelledConsumer {
	return CancelledConsumer{
		exchange: ex,
		out:      out,
		stop:     make(chan struct{}),
		ack:      make(chan types.Bytes),
		reject:   make(chan types.Bytes),
	}
}

func (cc *CancelledConsumer) StartConsuming() error {

	sink := make(chan *exchange.ExchangeInterfaceCancelled)

	sub, err := cc.exchange.WatchCancelled(nil, sink, make([][32]byte, 0))
	if err != nil {
		return err
	}

	cc.sub = sub

	go cc.consume(sink)
	go cc.logProcess()

	return nil
}

func (cc *CancelledConsumer) StopConsuming() {
	cc.sub.Unsubscribe()
	close(cc.stop)
}

func (cc *CancelledConsumer) consume(sink <-chan *exchange.ExchangeInterfaceCancelled) {
	for {
		select {
		case cancelled := <-sink:

			hash, err := types.NewBytes(string(cancelled.Hash[:]))
			if err != nil {
				log.Printf("hash err: %s", err)
				continue
			}

			cc.out <- NewCancelledMessage(*hash, cc.ack, cc.reject)
		case <-cc.stop:
			return
		}
	}
}

func (cc *CancelledConsumer) logProcess() {
	for {
		select {
		case tx := <-cc.reject:
			log.Printf("rejected cancel: %s", tx)
		case tx := <-cc.ack:
			log.Printf("ack cancel: %s", tx)
		case <-cc.stop:
			return
		}
	}
}
