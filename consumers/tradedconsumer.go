package consumers

import (
	"sync"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

type Block struct {
	Hash      common.Hash
	Timestamp types.Int
}

type TradedConsumer struct {
	sync.Mutex

	exchange *exchange.ExchangeInterface
	conn     *ethclient.Client

	out  chan<- *TradedMessage
	stop chan struct{}

	// @todo these 2 scenarios need handling
	ack    chan types.Bytes
	reject chan types.Bytes

	sub   event.Subscription
	block Block
}

func NewTradedConsumer(ex *exchange.ExchangeInterface, conn *ethclient.Client, out chan<- *TradedMessage) TradedConsumer {
	return TradedConsumer{
		exchange: ex,
		conn:     conn,
		out:      out,
		stop:     make(chan struct{}),
		ack:      make(chan types.Bytes),
		reject:   make(chan types.Bytes),
	}
}

func (tc *TradedConsumer) StartConsuming() error {

	sink := make(chan *exchange.ExchangeInterfaceTraded)

	sub, err := tc.exchange.WatchTraded(nil, sink, make([][32]byte, 0))
	if err != nil {
		return err
	}

	tc.sub = sub

	go tc.consume(sink)

	return nil
}

func (tc *TradedConsumer) StopConsuming() {
	tc.sub.Unsubscribe()
	close(tc.stop)
}

func (tc *TradedConsumer) consume(sink <-chan *exchange.ExchangeInterfaceTraded) {
	for {
		select {
		case trade := <-sink:
			tc.handleTrade(trade)
		case <-tc.stop:
			return
		}
	}

}

func (tc *TradedConsumer) blockTimestamp(hash common.Hash) (*types.Int, error) {
	tc.Lock()
	defer tc.Unlock()

	if tc.block.Hash != hash {
		h, err := tc.conn.HeaderByHash(nil, hash)
		if err != nil {
			return nil, err
		}

		tc.block.Hash = hash
		tc.block.Timestamp = types.Int{Int: *h.Time}
	}

	return &tc.block.Timestamp, nil
}

func (tc *TradedConsumer) handleTrade(trade *exchange.ExchangeInterfaceTraded) {
	time, err := tc.blockTimestamp(trade.Raw.BlockHash)
	if err != nil {
		// @todo think about how we can handle this gracefully
	}

	tc.out <- NewTradedMessage(types.NewTransaction(*trade, *time), tc.ack, tc.reject)
}
