package history

import (
	"github.com/DexyProject/dexy-go/types"
)

type History interface {
	GetHistory(pair types.Pair, user *types.Address, limit int) []types.Transaction
	InsertTransaction(transaction types.Transaction) error
}
