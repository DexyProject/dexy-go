package markets

import "github.com/DexyProject/dexy-go/types"

type Markets interface {
	InsertMarkets(markets []types.Market) error
	GetMarkets(tokens []types.Address) ([]types.Market, error)
}
