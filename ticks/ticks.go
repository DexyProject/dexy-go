package ticks

import (
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
)

type TickQuery interface {
	InsertTick(NewTick types.Transaction) error
	FetchTicks()
}
