package ticks

import (
	"github.com/DexyProject/dexy-go/types"
)

type Ticks interface {
	InsertTick(NewTick types.Tick) error
	FetchTicks(block int64) ([]types.Tick, error)
}
