package orderbook

import (
	"github.com/DexyProject/dexy-go/types"
)

type OrderBook interface {
	InsertOrder(order types.Order) error
	RemoveOrder(hash string) bool
	Bids(token types.Address, user *types.Address, limit int) []types.Order
	Asks(token types.Address, user *types.Address, limit int) []types.Order
	GetOrderByHash(hash string) *types.Order
}
