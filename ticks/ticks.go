package ticks

import (
	"github.com/DexyProject/dexy-go/types"
)

type Ticks interface {
	InsertTicks(ticks []types.Tick) error
	FetchTicks(token types.Address) ([]types.Tick, error)
	FetchAggregateVolumeForTokens(tokens []types.Address) (map[types.Address]types.Int, error)
	FetchLatestCloseForTokens(tokens []types.Address) (map[types.Address]float64, error)
}
