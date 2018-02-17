package consumers

import (
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
	exchange *exchange.ExchangeInterface
	conn     *ethclient.Client

	out chan<- types.Transaction

	sub   event.Subscription
	block Block
}

func NewTradedConsumer(ex *exchange.ExchangeInterface, conn *ethclient.Client, out chan<- types.Transaction) TradedConsumer {
	return TradedConsumer{exchange: ex, conn: conn, out: out}
}

func (tc *TradedConsumer) StartConsuming() error {

	sink := make(chan *exchange.ExchangeInterfaceTraded)

	sub, err := tc.exchange.WatchTraded(nil, sink, make([][32]byte, 0))
	if err != nil {
		return err // @todo better
	}

	tc.sub = sub

	go tc.consume(sink)

	return nil
}

func (tc *TradedConsumer) StopConsuming() {
	tc.sub.Unsubscribe()
}

func (tc *TradedConsumer) consume(sink <-chan *exchange.ExchangeInterfaceTraded) {
	for {

		trade := <-sink

		time, err := tc.blockTimestamp(trade.Raw.BlockHash)
		if err != nil {
			// @todo think about how we can handle this gracefully
		}

		tc.out <- types.NewTransaction(*trade, *time)
	}

}

func (tc *TradedConsumer) blockTimestamp(hash common.Hash) (*types.Int, error) {
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
