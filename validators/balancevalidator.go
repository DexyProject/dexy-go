package validators

import (
	"fmt"
	"strconv"

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
	exchangeInterface, err := exchange.NewExchangeInterface(o.Exchange, balanceSession.conn)

	if err != nil {
		return false, fmt.Errorf("could not connect to contract session")
	}

	balance, err := exchangeInterface.BalanceOf(nil, o.Give.Token, o.User)
	if err != nil {
		return false, fmt.Errorf("could not get balance from contract")
	}

	balanceFloat := float64(balance.Int64())
	giveAmount, err := strconv.ParseFloat(o.Give.Amount, 64)

	if err != nil {
		return false, fmt.Errorf("error parsing o.give.amount")
	}

	return balanceFloat >= giveAmount, nil
}
