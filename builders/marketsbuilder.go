package builders

import (
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
)

type MarketsBuilder struct {

}

// @todo other type for bid/ask
func (mb *MarketsBuilder) build(ticks []ticks.Ticks) (types.Market) {
	return nil // @todo
}