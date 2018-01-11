package validators

import (
	"github.com/DexyProject/dexy-go/exchange"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"fmt"
)

type BalanceValidator interface {
	CheckBalance(tokenAddr, userAddr common.Address) (*big.Int, error)
}

type BalanceValidatorSession struct {
	conn bind.ContractBackend

}
func (balanceSession *BalanceValidatorSession) CheckBalance(tokenAddr, userAddr common.Address) (*big.Int, error) {
	exchange, err := exchange.NewExchangeInterface(tokenAddr, balanceSession.conn)//Can pass nil for CallOpts and TransactOpts
	if err != nil {
		return nil, fmt.Errorf("could not connect to contract session")
	}

	return exchange.BalanceOf(nil, tokenAddr, userAddr)
}