package history

import (
	"fmt"
	"math/big"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type HistoryAggregation struct {
	connection string
	session    *mgo.Session
}

func NewHistoryAggregation(connection string) (*HistoryAggregation, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &HistoryAggregation{connection: connection, session: session}, nil
}

func (history *HistoryAggregation) AggregateTransactions(block int64, transactions []types.Transaction) ([]types.Tick, error) {
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
		prices, txindex := getPrices(mappedTokens[token])
		openIndex, closeIndex := calcOpenCloseIndex(mappedTokens[token])
		openPrice, closePrice := calcOpenClosePrice(prices, txindex, openIndex, closeIndex)
		high, low := calcHighLow(prices)

		ticks = append(
			ticks, types.Tick{
				Pair:      pair,
				Block:     block,
				Volume:    types.Int{*volume},
				Open:      openPrice,
				Close:     closePrice,
				High:      high,
				Low:       low,
				Timestamp: mappedTokens[token][1].Timestamp,
			},
		)
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

func calcHighLow(prices []float64) (float64, float64) {
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

func getPrices(transactions []types.Transaction) ([]float64, []uint) {
	var prices []float64
	var txindex []uint
	for _, tt := range transactions {
		newPrice, _ := tt.Get.CalcPrice(tt.Give, types.HexToAddress(types.ETH_ADDRESS))
		prices = append(prices, newPrice)
		txindex = append(txindex, tt.TransactionIndex)
	}

	return prices, txindex
}

func calcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
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

func calcOpenClosePrice(prices []float64, txindex []uint, OpenIndex, CloseIndex uint) (float64, float64) {
	var openPrice, closePrice float64
	for i, tt := range txindex {
		switch tt {
		case OpenIndex:
			openPrice = prices[i]
		case CloseIndex:
			closePrice = prices[i]
		}
	}

	return openPrice, closePrice
}

func getPair(token types.Address) types.Pair {
	return types.Pair{types.HexToAddress(types.ETH_ADDRESS), token}
}

func groupTokens(transactions []types.Transaction) map[types.Address][]types.Transaction {
	m := make(map[types.Address][]types.Transaction)
	for _, t := range transactions {
		if t.Get.Token == types.HexToAddress(types.ETH_ADDRESS) {
			m[t.Give.Token] = append(m[t.Give.Token], t)
		} else {
			m[t.Get.Token] = append(m[t.Get.Token], t)
		}
	}

	return m
}
