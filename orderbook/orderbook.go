package orderbook

import (
	"github.com/DexyProject/dexy-go/types"
)

type OrderBook interface {
	InsertOrder(order types.Order) error
	RemoveOrder(hash types.Hash) bool
	Bids(token types.Address, limit int) []types.Order
	Asks(token types.Address, limit int) []types.Order
	UpdateOrderFilledAmount(hash types.Hash, amount types.Int) error
	GetOrderByHash(hash types.Hash) *types.Order
	GetOrders(token types.Address, user *types.Address, limit int) []types.Order
	GetMarkets(tokens []types.Address) []types.Market
}
