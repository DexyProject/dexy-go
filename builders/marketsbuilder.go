package builders

import (
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
)

type MarketsBuilder struct {
}

// @todo other type for bid/ask
// @todo ticks mapping
func (mb *MarketsBuilder) Build(ticks []types.Tick, asks types.Prices, bids types.Prices) []types.Market {
	return nil // @todo
}
