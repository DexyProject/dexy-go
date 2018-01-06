package history

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/DexyProject/dexy-go/types"
)

type History interface {
	GetHistory(token common.Address) ([]types.Transaction)
}
