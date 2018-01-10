package validators

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"fmt"
)

type BalanceValidator struct {
	contract *exchange.ExchangeInterface
	callOpts bind.CallOpts
	transactOpts bind.TransactOpts

}
func (contractSession BalanceValidator) CheckBalance(token common.Address, user common.Address) (*big.Int, error) {
	balance, err:= exchange.ExchangeInterfaceSession{Contract:contractSession.contract, CallOpts:contractSession.callOpts,
		TransactOpts:contractSession.transactOpts}.BalanceOf(token, user) //Can pass nil for CallOpts and TransactOpts
	if err != nil {
		return nil, fmt.Errorf("could not connect to contract session")
	}
	return balance, nil
}


