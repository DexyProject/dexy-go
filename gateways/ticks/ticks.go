package ticks

import (
	"github.com/DexyProject/dexy-go/types"
)

type Ticks interface {
	InsertTicks(ticks []types.Tick) error
	FetchTicks(token types.Address) ([]types.Tick, error)
}
