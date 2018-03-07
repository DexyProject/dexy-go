package repositories

import (
	"fmt"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenRepository interface {
	GetDecimals(token types.Address) (uint8, error)
}

type CacheTokensRepository struct {
	connection *ethclient.Client
	decimals   map[types.Address]uint8
}

func NewCacheTokensRepository(connection string) (*CacheTokensRepository, error) {
	conn, err := ethclient.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not dial ethclient")
	}

	return &CacheTokensRepository{
		connection: conn,
		decimals:   make(map[types.Address]uint8),
	}, nil
}

func (repository *CacheTokensRepository) GetDecimals(token types.Address) (uint8, error) {
	if _, ok := repository.decimals[token]; ok {
		return repository.decimals[token], nil
	}

	erc20, err := contracts.NewERC20(token.Address, repository.connection)
	if err != nil {
		return 0, fmt.Errorf("could not access contract for token")
	}

	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return 0, fmt.Errorf("could not get decimals() from contract")
	}

	repository.decimals[token] = decimals
	return decimals, nil
}
