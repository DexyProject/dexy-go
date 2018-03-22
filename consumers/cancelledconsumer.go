package consumers

import (
	"log"
	"sync"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/event"
)

type CancelledConsumer struct {
	sync.Mutex

	exchange *contracts.Exchange

	out  chan<- *CancelledMessage
	stop chan struct{}

	ack    chan types.Hash
	reject chan types.Hash

	sub event.Subscription
}

func NewCancelledConsumer(ex *contracts.Exchange, out chan<- *CancelledMessage) CancelledConsumer {
	return CancelledConsumer{
		exchange: ex,
		out:      out,
		stop:     make(chan struct{}),
		ack:      make(chan types.Hash),
		reject:   make(chan types.Hash),
	}
}

func (cc *CancelledConsumer) StartConsuming() error {

	sink := make(chan *contracts.ExchangeCancelled)

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

func (cc *CancelledConsumer) consume(sink <-chan *contracts.ExchangeCancelled) {
	for {
		select {
		case cancelled := <-sink:
			hash := types.Hash{}
			hash.SetBytes(cancelled.Hash[:])

			cc.out <- NewCancelledMessage(hash, cc.ack, cc.reject)
		case <-cc.stop:
			return
		}
	}
}

func (cc *CancelledConsumer) logProcess() {
	for {
		select {
		case tx := <-cc.reject:
			log.Printf("rejected cancel: %s", tx.String())
		case tx := <-cc.ack:
			log.Printf("ack cancel: %s", tx.String())
		case <-cc.stop:
			return
		}
	}
}
