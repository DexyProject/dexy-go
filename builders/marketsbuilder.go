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

func NewMarketsBuilder(repo repositories.TokenRepository) MarketsBuilder {
	return MarketsBuilder{repo: repo}
}

func (mb *MarketsBuilder) Build(tokens []types.Address, ticks Ticks, asks types.Prices, bids types.Prices) []types.Market {

	markets := make([]types.Market, 0)

	for _, token := range tokens {

		market := types.Market{}

		market.Token = token

		if tick, ok := ticks[token]; ok {
			market.Last = tick.Close

			vol, _ := normalize(tick.Volume.String(), 18.0)
			market.Volume = vol
		}

		decimals, err := mb.repo.GetDecimals(token)
		if err != nil {
			log.Error("failed loading decimals", zap.String("error", token.String()), zap.Error(err))
			continue
		}

		market.Ask = getPrice(token, asks, decimals)
		market.Bid = getPrice(token, bids, decimals)

		//if (types.Market{}) == market {
		//	continue
		//}

		markets = append(markets, market)
	}

	return markets
}

func getPrice(token types.Address, prices types.Prices, decimals uint8) (float64) {
	price, ok := prices[token]
	if !ok {
		return 0.0
	}

	p, err := calculatePrice(price.Quote, price.Base, decimals)
	if err != nil {
		log.Error("failed to calculate price", zap.Error(err))
	}

	return p
}

func calculatePrice(quote string, base string, decimals uint8) (float64, error) {
	q, err := normalize(quote, float64(decimals))
	if err != nil {
		return 0.0, fmt.Errorf("failed to create float from quote %s", quote)
	}

	b, err := normalize(base, 18.0)
	if err != nil {
		return 0.0, fmt.Errorf("failed to create float from base %s", base)
	}

	return (b / math.Pow(10.0, 18.0)) / (q / math.Pow(10.0, float64(decimals))), nil
}

func normalize(number string, pow float64) (float64, error) {
	bf, ok := new(big.Float).SetString(number)
	if !ok {
		return 0.0, fmt.Errorf("failed to create float from %s", number)
	}

	f, _ := bf.Float64()

	return f / math.Pow(10.0, pow), nil
}
