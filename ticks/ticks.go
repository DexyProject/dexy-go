package ticks

import (
	"github.com/DexyProject/dexy-go/types"
)

type Ticks interface {
	InsertTick(NewTick types.Tick) error
	FetchTicks(token types.Address) ([]types.Tick, error)
}
