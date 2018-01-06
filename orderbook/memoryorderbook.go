package orderbook

import (
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/DexyProject/dexy-go/types"
)

type MemoryOrderBook struct {
	orders []types.Order
}

func (ob *MemoryOrderBook) InsertOrder(NewOrder types.Order) error {
	// Check if hash exists, then verify signature to add to OB
	if ob.GetOrderByHash(NewOrder.Hash) != nil {
		return fmt.Errorf("order exists in orderbook")
	}

	hash, err := NewOrder.OrderHash()
	if err != nil {
		return fmt.Errorf("could not create order hash")
	}

	if !NewOrder.Signature.Verify(NewOrder.User, hash) {
		return fmt.Errorf("signature could not be verified")
	}

	ob.orders = append(ob.orders, NewOrder)
	return nil
}

func (ob *MemoryOrderBook) RemoveOrder(hash string) bool {
	for i := range ob.orders {
		if ob.orders[i].Hash == hash {
			ob.orders = append(ob.orders[:i], ob.orders[i+1:]...)
			return true
		}
	}

	return false
}

func (ob *MemoryOrderBook) Bids(token common.Address, user *common.Address, limit int) []types.Order {

	orders := []types.Order{}

	for i := range ob.orders {
		if user != nil && ob.orders[i].User != *user {
			continue
		}

		if ob.orders[i].Get.Token == token {
			orders = append(orders, ob.orders[i])
		}
	}

	sort.Slice(orders[:], func(i, j int) bool {
		return orders[i].Price > orders[j].Price // sorting bids in descending order
	})

	if len(orders) <= limit {
		return orders
	}

	return orders[0 : limit-1]
}

func (ob *MemoryOrderBook) Asks(token common.Address, user *common.Address, limit int) []types.Order {

	orders := []types.Order{}

	for i := range ob.orders {
		if user != nil && ob.orders[i].User != *user {
			continue
		}

		if ob.orders[i].Give.Token == token {
			orders = append(orders, ob.orders[i])
		}
	}

	sort.Slice(orders[:], func(i, j int) bool {
		return orders[i].Price < orders[j].Price // sorting bids in descending order
	})

	if len(orders) <= limit {
		return orders
	}

	return orders[0 : limit-1]
}

func (ob *MemoryOrderBook) GetOrderByHash(hash string) *types.Order {
	for i := range ob.orders {
		if ob.orders[i].Hash == hash {
			return &ob.orders[i]
		}
	}

	return nil
}
