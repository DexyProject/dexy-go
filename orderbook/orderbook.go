package orderbook

import (
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type OrderBook interface {
	InsertOrder(order types.Order) error
	RemoveOrder(hash string) bool
	Bids(token common.Address, user *common.Address, limit int) []types.Order
	Asks(token common.Address, user *common.Address, limit int) []types.Order
	GetOrderByHash(hash string) *types.Order
}
