package consumers

import (
	"context"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"go.uber.org/zap"
)

type TradedConsumer struct {
	exchange *contracts.Exchange

	conn *ethclient.Client

	out  chan<- *TradedMessage
	stop chan struct{}

	ack    chan types.Hash
	reject chan types.Hash

	sub event.Subscription

	blocks map[common.Hash]types.Int
}

func NewTradedConsumer(ex *contracts.Exchange, conn *ethclient.Client, out chan<- *TradedMessage) TradedConsumer {
	return TradedConsumer{
		exchange: ex,
		conn:     conn,
		out:      out,
		stop:     make(chan struct{}),
		ack:      make(chan types.Hash),
		reject:   make(chan types.Hash),
		blocks:   make(map[common.Hash]types.Int),
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
			log.Debug("rejected trade", zap.String("tx", tx.String()))
		case tx := <-tc.ack:
			log.Debug("ack trade", zap.String("tx", tx.String()))
		case <-tc.stop:
			return
		}
	}
}

func (tc *TradedConsumer) blockTimestamp(hash common.Hash) (*types.Int, error) {
	_, ok := tc.blocks[hash]
	if !ok {
		h, err := tc.conn.HeaderByHash(context.Background(), hash)
		if err != nil {
			return nil, err
		}

		tc.blocks[hash] = types.Int{Int: *h.Time}
	}

	b := tc.blocks[hash]
	return &b, nil
}

func (tc *TradedConsumer) handleTrade(trade *contracts.ExchangeTraded) {
	time, err := tc.blockTimestamp(trade.Raw.BlockHash)
	if err != nil {
		// @todo think about how we can handle this gracefully
		panic(err)
	}

	tc.out <- NewTradedMessage(types.NewTransaction(*trade, *time), tc.ack, tc.reject)
}
