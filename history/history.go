package history

import (
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type History interface {
	GetHistory(token common.Address, user *common.Address, limit int) []types.Transaction
}
