package markets

import "github.com/DexyProject/dexy-go/types"

type Markets interface {
	GetMarkets([]types.Address) ([]types.Market, error)
}
