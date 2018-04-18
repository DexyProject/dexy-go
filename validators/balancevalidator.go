package validators

import (
	"fmt"
	"math/big"

	"github.com/DexyProject/dexy-go/balances"
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
)

type BalanceValidator interface {
	CheckBalance(o types.Order) (bool, error)
}

type RPCBalanceValidator struct {
	vault    *contracts.Vault
	balances balances.Balances
}

func NewRPCBalanceValidator(vault *contracts.Vault, balances balances.Balances) *RPCBalanceValidator {
	return &RPCBalanceValidator{
		vault:    vault,
		balances: balances,
	}
}

func (b *RPCBalanceValidator) CheckBalance(o types.Order) (bool, error) {
	balance, err := b.vault.BalanceOf(nil, o.Make.Token.Address, o.User.Address)
	if err != nil {
		return false, fmt.Errorf("could not get balance from contract: %s", err.Error())
	}

	if balance.Sign() == 0 {
		return false, nil
	}

	onOrders, err := b.balances.OnOrders(o.User, o.Make.Token)
	if err != nil {
		return false, fmt.Errorf("balances error: %v", err.Error())
	}

	// (balances - onOrders) >= amount
	return new(big.Int).Sub(balance, &onOrders.Int).Cmp(&o.Make.Amount.Int) >= 0, nil
}
