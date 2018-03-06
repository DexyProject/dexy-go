package history

import (
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
)

type TokensRepository struct {
	decimals   map[types.Address]uint8
}

func NewTokensRepository() *TokensRepository {
	return &TokensRepository{decimals : make(map[types.Address]uint8)}
}

func (repository *TokensRepository) Decimals(token types.Address) (uint8, error) {
	if _, ok := repository.decimals[token]; ok {
		return repository.decimals[token], nil
	}
	erc20, err := contracts.NewERC20(token.Address, nil)
	if err != nil {
		return 0, fmt.Errorf("could not access contract for token")
	}
	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return 0, fmt.Errorf("could not get decimals() from contract")
	}
	repository.decimals[token] = decimals;
	return decimals, nil
}