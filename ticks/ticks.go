package ticks

import (
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
)

type Ticks interface {
	InsertTick(NewTick types.Tick) error
	FetchTicks(block int64) ([]types.Tick, error)
}
