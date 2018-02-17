package validators

import (
	"fmt"
	"math/big"

	"github.com/DexyProject/dexy-go/balances"
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type BalanceValidator interface {
	CheckBalance(o types.Order) (bool, error)
}

type RPCBalanceValidator struct {
	Conn     bind.ContractBackend
	Balances balances.Balances
}

func (balanceSession *RPCBalanceValidator) CheckBalance(o types.Order) (bool, error) {
	exchangeInterface, err := exchange.NewExchangeInterfaceCaller(o.Exchange.Address, balanceSession.Conn)

	if err != nil {
		return false, fmt.Errorf("could not connect to contract session")
	}

	balance, err := exchangeInterface.BalanceOf(nil, o.Give.Token.Address, o.User.Address)
	if err != nil {
		return false, fmt.Errorf("could not get balance from contract")
	}

	onOrders, err := balanceSession.Balances.OnOrders(o.User, o.Give.Token)
	if err != nil {
		return false, fmt.Errorf("balances error: %v", err.Error())
	}

	if balance.String() == "0" {
		return false, nil
	}

	// (balances - onOrders) >= amount
	return new(big.Int).Sub(balance, &onOrders.Int).Cmp(&o.Give.Amount.Int) >= 0, nil
}
