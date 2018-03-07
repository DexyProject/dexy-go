package repositories

import (
	"github.com/DexyProject/dexy-go/types"
)

type MockCacheTokensRepository struct {
	decimals map[types.Address]uint8
}

func (m *MockCacheTokensRepository) GetDecimals(token types.Address) (uint8, error) {
	return m.decimals[token], nil
}

func (m *MockCacheTokensRepository) AddToken(token types.Address, decimals uint8) {
	m.decimals[token] = decimals
}
