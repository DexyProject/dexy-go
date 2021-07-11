package history

import (
	"github.com/DexyProject/dexy-go/types"
)

type History interface {
	GetHistory(token types.Address, user *types.Address, limit int) []types.Transaction
	InsertTransaction(transaction types.Transaction) error
	GetTransactionsInBlock(block uint64) ([]types.Transaction, error)
}
