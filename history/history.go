package history

import (
	"github.com/DexyProject/dexy-go/types"
)

type History interface {
	GetHistory(quote types.Address, base types.Address, user *types.Address, limit int) []types.Transaction
	InsertTransaction(transaction types.Transaction) error
}
