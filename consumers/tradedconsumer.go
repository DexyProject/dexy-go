package consumers

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

type Block struct {
	Hash      common.Hash
	Timestamp types.Int
}

type TradedConsumer struct {
	Exchange exchange.ExchangeInterface

	sub   event.Subscription
	block Block

	out chan<- types.Transaction
}

func (tc *TradedConsumer) StartConsuming() error {

	sink := make(chan *exchange.ExchangeInterfaceTraded)

	sub, err := tc.Exchange.WatchTraded(nil, sink, make([][32]byte, 0))
	if err != nil {
		// @todo return
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
	if tc.block.Hash == hash {
		return &tc.block.Timestamp, nil
	}

	// @todo query time and update

	return nil, nil // @todo
}
