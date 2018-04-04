package builders

import (
	"fmt"
	"math"
	"math/big"

	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/types"
	"go.uber.org/zap"
)

type Ticks map[types.Address]types.Tick

type MarketsBuilder struct {
	repo repositories.TokenRepository
}

func (mb *MarketsBuilder) Build(tokens []types.Address, ticks Ticks, asks types.Prices, bids types.Prices) []types.Market {

	markets := make([]types.Market, 0)

	for _, token := range tokens {

		market := types.Market{}

		if tick, ok := ticks[token]; ok {
			market.Last = tick.Close
			market.Volume = tick.Volume
		}

		decimals, err := mb.repo.GetDecimals(token)
		if err != nil {
			log.Error("failed loading decimals", zap.Error(err))
			continue
		}

		if ask, ok := asks[token]; ok {
			p, err := calculatePrice(ask.Quote, ask.Base, decimals)
			if err != nil {
				log.Error("failed to calculate price", zap.Error(err))
			}

			market.Ask = p
		}

		if bid, ok := bids[token]; ok {
			p, err := calculatePrice(bid.Quote, bid.Base, decimals)
			if err != nil {
				log.Error("failed to calculate price", zap.Error(err))
			}

			market.Bid = p
		}

		//if (types.Market{}) == market {
		//	continue
		//}

		markets = append(markets, market)
	}

	return markets
}

func calculatePrice(quote string, base string, decimals uint8) (float64, error) {
	q, ok := new(big.Float).SetString(quote)
	if !ok {
		return 0.0, fmt.Errorf("failed to create float from quote %s", quote)
	}

	b, ok := new(big.Float).SetString(base)
	if !ok {
		return 0.0, fmt.Errorf("failed to create float from base %s", base)
	}

	bf, _ := b.Float64()
	qf, _ := q.Float64()

	return (bf / math.Pow(10.0, 18.0)) / (qf / math.Pow(10.0, float64(decimals))), nil
}
