package validators

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"github.com/DexyProject/dexy-go/types"
	"strconv"
)

type BalanceValidator interface {
	CheckBalance(o types.Order) error
}

type BalanceValidatorSession struct {
	conn bind.ContractBackend
}

func (balanceSession *BalanceValidatorSession) CheckBalance(o types.Order) error {
	exchangeInterface, err := exchange.NewExchangeInterface(o.Exchange, balanceSession.conn)

	if err != nil {
		return fmt.Errorf("could not connect to contract session")
	}

	balance, err := exchangeInterface.BalanceOf(nil, o.Exchange, o.User)
	if err != nil {
		return fmt.Errorf("could not get balance from contract")
	}

	balanceFloat := float64(balance.Int64()) // nasty type conversions
	giveAmount, err := strconv.ParseFloat(o.Give.Amount, 64)

	if err != nil {
		return fmt.Errorf("error parsing o.give.amount")
	}
	if balanceFloat < giveAmount {
		return fmt.Errorf("balance too low")
	}
	return nil
}
