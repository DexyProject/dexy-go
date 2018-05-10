package builders

import (
	"fmt"
	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/math"
)

type HistoryBuilder struct {
	repository repositories.TokenRepository
}

func (hb *HistoryBuilder) build(transactions []types.Transaction) ([]types.Tick, error) {
	ticks := make([]types.Tick, 0)

	mappedTokens := hb.groupTokens(transactions)

	block := transactions[0].BlockNumber

	for token := range mappedTokens {
		decimals, err := hb.repository.GetDecimals(token)
		if err != nil {
			return nil, err
		}
		pair := hb.getPair(token)
		volume := hb.calcVolume(mappedTokens[token])
		prices, txindex := hb.getPrices(mappedTokens[token], decimals)
		openIndex, closeIndex := hb.calcOpenCloseIndex(mappedTokens[token])
		openPrice, closePrice := hb.calcOpenClosePrice(prices, txindex, openIndex, closeIndex)
		high, low := hb.calcHighLow(prices)

		ticks = append(
			ticks, types.Tick{
				Pair:      pair,
				Block:     block,
				Volume:    *volume,
				Open:      openPrice,
				Close:     closePrice,
				High:      high,
				Low:       low,
				Timestamp: mappedTokens[token][0].Timestamp,
			},
		)
	}

	return ticks, nil
}

func (hb *HistoryBuilder) calcVolume(transactions []types.Transaction) *types.Int {
	volume := new(types.Int)
	for _, tt := range transactions {
		switch types.HexToAddress(types.ETH_ADDRESS) {
		case tt.Make.Token:
			volume.Add(&volume.Int, &tt.Make.Amount.Int)
		case tt.Take.Token:
			volume.Add(&volume.Int, &tt.Take.Amount.Int)
		}
	}

	return volume
}

func (hb *HistoryBuilder) calcHighLow(prices []float64) (float64, float64) {
	high, low := prices[0], prices[0]
	for _, p := range prices {
		if high < p {
			high = p
		}
		if low > p {
			low = p
		}
	}

	return high, low
}

func (hb *HistoryBuilder) calcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
	openIndex, closeIndex := transactions[0].TransactionIndex, transactions[0].TransactionIndex
	for _, tt := range transactions {
		if openIndex > tt.TransactionIndex {
			openIndex = tt.TransactionIndex
		}
		if closeIndex < tt.TransactionIndex {
			closeIndex = tt.TransactionIndex
		}
	}

	return openIndex, closeIndex
}

func (hb *HistoryBuilder) calcOpenClosePrice(prices []float64, txindex []uint, openIndex, closeIndex uint) (float64, float64) {
	var openPrice, closePrice float64
	for i, tt := range txindex {
		if tt == openIndex {
			openPrice = prices[i]
		}

		if tt == closeIndex {
			closePrice = prices[i]
		}
	}

	return openPrice, closePrice
}

func (hb *HistoryBuilder) getPair(token types.Address) types.Pair {
	return types.Pair{Quote: token, Base: types.HexToAddress(types.ETH_ADDRESS)}
}

func (hb *HistoryBuilder) groupTokens(transactions []types.Transaction) map[types.Address][]types.Transaction {
	m := make(map[types.Address][]types.Transaction)
	for _, t := range transactions {
		if t.Take.Token == types.HexToAddress(types.ETH_ADDRESS) {
			m[t.Make.Token] = append(m[t.Make.Token], t)
		} else {
			m[t.Take.Token] = append(m[t.Take.Token], t)
		}
	}

	return m
}

func (hb *HistoryBuilder) calcPrice(t types.Transaction, base types.Address, decimals uint8) (float64, error) {
	if t.Make.Amount.Sign() <= 0 || t.Take.Amount.Sign() <= 0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}

	baseAmount := t.Make.Amount
	quoteAmount := t.Take.Amount

	if t.Take.Token == base {
		baseAmount = t.Take.Amount
		quoteAmount = t.Make.Amount
	}

	return math.ToUnitAmount(baseAmount, 18.0) / math.ToUnitAmount(quoteAmount, decimals), nil
}

func (hb *HistoryBuilder) getPrices(transactions []types.Transaction, decimals uint8) ([]float64, []uint) {
	prices := make([]float64, 0)
	txindex := make([]uint, 0)
	for _, tt := range transactions {
		newPrice, _ := hb.calcPrice(tt, types.HexToAddress(types.ETH_ADDRESS), decimals)
		prices = append(prices, newPrice)
		txindex = append(txindex, tt.TransactionIndex)
	}

	return prices, txindex
}
