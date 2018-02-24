package history

import (
	"github.com/DexyProject/dexy-go/types"
	"math/big"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func (history *MongoHistory) AggregateTransactions(block int64, transactions []types.Transaction) ([]types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()
	c := session.DB(DBName).C(FileName)

	var ticks []types.Tick
	var matchedTransactions []types.Transaction

	matchBlock := bson.M{"$match": bson.M{"$transactions.block": block}}

	err := c.Pipe([]bson.M{matchBlock}).All(&matchedTransactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}

	mappedTokens := groupTokens(matchedTransactions)
	for token := range mappedTokens {
		pair := getPair(token)
		volume := calcVolume(mappedTokens[token])
		prices := getPrices(mappedTokens[token])
		openIndex, closeIndex := calcOpenCloseIndex(mappedTokens[token])
		openPrice, closePrice := calcOpenClosePrice(prices, openIndex, closeIndex)
		high, low := calcHighLow(prices)

		ticks = append(ticks,
			types.Tick{
				Pair: pair,
				Block: block,
				Volume: types.Int{*volume},
				Open: openPrice,
				Close: closePrice,
				High: high,
				Low: low,
			})
	}

	return ticks, nil
}

func calcVolume(transactions []types.Transaction) *big.Int {
	volume := new(big.Int)
	for _, tt := range transactions {
		switch {
		case tt.Give.Token != types.HexToAddress(types.ETH_ADDRESS):
			volume.Add(volume, &tt.Give.Amount.Int)
		case tt.Get.Token != types.HexToAddress(types.ETH_ADDRESS):
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}

	return volume
}

func calcHighLow(prices []types.Price) (float64, float64) {
	high, low := prices[0].Price, prices[0].Price
	for _, p := range prices {
		switch {
		case high < p.Price:
			high = p.Price
		case low > p.Price:
			low = p.Price
		}
	}

	return high, low
}

func getPrices(transactions []types.Transaction) []types.Price {
	var prices []types.Price
	for _, tt := range transactions {
		newPrice, _ := tt.Get.CalcPrice(tt.Give, types.HexToAddress(types.ETH_ADDRESS))
		priceStruct := types.Price{tt.TransactionIndex, newPrice}
		prices = append(prices, priceStruct)
	}

	return prices
}

func calcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
	openIndex, closeIndex := transactions[0].TransactionIndex, transactions[0].TransactionIndex
	for _, tt := range transactions {
		switch {
		case openIndex > tt.TransactionIndex:
				openIndex = tt.TransactionIndex
		case closeIndex < tt.TransactionIndex:
				closeIndex = tt.TransactionIndex
		}
	}

	return openIndex, closeIndex
}

func calcOpenClosePrice(prices []types.Price, OpenIndex, CloseIndex uint) (float64, float64) {
	var openPrice, closePrice float64
	for _, tt := range prices {
		switch {
		case tt.TransactionIndex == OpenIndex:
				openPrice = tt.Price
		case tt.TransactionIndex == CloseIndex:
			closePrice = tt.Price
		}
	}

	return openPrice, closePrice
}

func getPair(token types.Address) types.Pair {
	newPair := types.Pair{types.HexToAddress(types.ETH_ADDRESS), token}
	return newPair
}

func groupTokens(transactions []types.Transaction) map[types.Address][]types.Transaction{
	var m map[types.Address][]types.Transaction
	m = make(map[types.Address][]types.Transaction)
	for _, t := range transactions {
		if t.Get.Token == types.HexToAddress(types.ETH_ADDRESS) {
			m[t.Give.Token] = append(m[t.Give.Token], t)
		} else {
			m[t.Get.Token] = append(m[t.Get.Token], t)
		}
	}

	return m
}
