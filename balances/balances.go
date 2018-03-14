package balances

import "github.com/DexyProject/dexy-go/types"

type Balances interface {
	OnOrders(user types.Address, token types.Address) (*types.Int, error)
	Underfunded(user types.Address, token types.Address) (*types.Int, error)
}
