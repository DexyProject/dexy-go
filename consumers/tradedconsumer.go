package consumers

import (
	"context"
	"log"
	"sync"

	"github.com/DexyProject/dexy-go/contracts"
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

	exchange *contracts.Exchange
	conn     *ethclient.Client

	out  chan<- *TradedMessage
	stop chan struct{}

	ack    chan types.Bytes
	reject chan types.Bytes

	sub   event.Subscription
	block Block
}

func NewTradedConsumer(ex *contracts.Exchange, conn *ethclient.Client, out chan<- *TradedMessage) TradedConsumer {
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

	sink := make(chan *contracts.ExchangeTraded)

	sub, err := tc.exchange.WatchTraded(nil, sink, make([][32]byte, 0))
	if err != nil {
		return err
	}

	tc.sub = sub

	go tc.consume(sink)
	go tc.logProcess()

	return nil
}

func (tc *TradedConsumer) StopConsuming() {
	tc.sub.Unsubscribe()
	close(tc.stop)
}

func (tc *TradedConsumer) consume(sink <-chan *contracts.ExchangeTraded) {
	for {
		select {
		case trade := <-sink:
			tc.handleTrade(trade)
		case <-tc.stop:
			return
		}
	}
}

func (tc *TradedConsumer) logProcess() {
	for {
		select {
		case tx := <-tc.reject:
			log.Printf("rejected tx: %s", tx)
		case tx := <-tc.ack:
			log.Printf("ack tx: %s", tx)
		case <-tc.stop:
			return
		}
	}
}

func (tc *TradedConsumer) blockTimestamp(hash common.Hash) (*types.Int, error) {
	tc.Lock()
	defer tc.Unlock()

	if tc.block.Hash != hash {
		h, err := tc.conn.HeaderByHash(context.Background(), hash)
		if err != nil {
			return nil, err
		}

		tc.block.Hash = hash
		tc.block.Timestamp = types.Int{Int: *h.Time}
	}

	return &tc.block.Timestamp, nil
}

func (tc *TradedConsumer) handleTrade(trade *contracts.ExchangeTraded) {
	time, err := tc.blockTimestamp(trade.Raw.BlockHash)
	if err != nil {
		// @todo think about how we can handle this gracefully
		panic(err)
	}

	tc.out <- NewTradedMessage(types.NewTransaction(*trade, *time), tc.ack, tc.reject)
}
