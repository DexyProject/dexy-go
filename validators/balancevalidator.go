package validators

import (
	"fmt"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type BalanceValidator interface {
	CheckBalance(o types.Order) (bool, error)
}

type BalanceValidatorSession struct {
	conn bind.ContractBackend
}

func (balanceSession *BalanceValidatorSession) CheckBalance(o types.Order) (bool, error) {
	exchangeInterface, err := exchange.NewExchangeInterface(o.Exchange.Address, balanceSession.conn)

	if err != nil {
		return false, fmt.Errorf("could not connect to contract session")
	}

	balance, err := exchangeInterface.BalanceOf(nil, o.Give.Token.Address, o.User.Address)
	if err != nil {
		return false, fmt.Errorf("could not get balance from contract")
	}

	return balance.Cmp(&o.Give.Amount.Int) >= 0, nil
}
