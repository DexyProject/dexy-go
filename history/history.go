package history

import (
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
)

type History interface {
	GetHistory(token types.Address, user *types.Address, limit int) []types.Transaction
	AggregateTransactions(block int) ([]bson.M, error)
}