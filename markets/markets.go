package markets

import "github.com/DexyProject/dexy-go/types"

type Markets interface {
	InsertMarkets([]types.Market) error
	GetMarkets([]types.Address) ([]types.Market, error)
}
