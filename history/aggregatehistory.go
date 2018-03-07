package history

import (
	"fmt"
	"math"
	"math/big"

	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type HistoryAggregation struct {
	connection string
	session    *mgo.Session
	repository repositories.TokenRepository
}

func NewHistoryAggregation(connection string, repository repositories.TokenRepository) (*HistoryAggregation, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &HistoryAggregation{
		connection: connection,
		session:    session,
		repository: repository,
	}, nil
}

func (history *HistoryAggregation) AggregateTransactions(block int64) ([]types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()
	c := session.DB(DBName).C(FileName)

	var ticks []types.Tick
	var transactions []types.Transaction

	matchBlock := bson.M{"$match": bson.M{"block": block}}

	err := c.Pipe([]bson.M{matchBlock}).All(&transactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions: %s", err)
	}

	mappedTokens := history.groupTokens(transactions)
	for token := range mappedTokens {
		decimals, err := history.repository.GetDecimals(token)
		if err != nil {
			return nil, err
		}
		pair := history.getPair(token)
		volume := history.calcVolume(mappedTokens[token])
		prices, txindex := history.getPrices(mappedTokens[token], decimals)
		openIndex, closeIndex := history.calcOpenCloseIndex(mappedTokens[token])
		openPrice, closePrice := history.calcOpenClosePrice(prices, txindex, openIndex, closeIndex)
		high, low := history.calcHighLow(prices)

		ticks = append(
			ticks, types.Tick{
				Pair:      pair,
				Block:     block,
				Volume:    types.Int{*volume},
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

func (history *HistoryAggregation) calcVolume(transactions []types.Transaction) *big.Int {
	volume := new(big.Int)
	for _, tt := range transactions {
		switch types.HexToAddress(types.ETH_ADDRESS) {
		case tt.Give.Token:
			volume.Add(volume, &tt.Give.Amount.Int)
		case tt.Get.Token:
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}

	return volume
}

func (history *HistoryAggregation) calcHighLow(prices []float64) (float64, float64) {
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

func (history *HistoryAggregation) calcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
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

func (history *HistoryAggregation) calcOpenClosePrice(prices []float64, txindex []uint, OpenIndex, CloseIndex uint) (float64, float64) {
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

func (history *HistoryAggregation) getPair(token types.Address) types.Pair {
	return types.Pair{token, types.HexToAddress(types.ETH_ADDRESS)}
}

func (history *HistoryAggregation) groupTokens(transactions []types.Transaction) map[types.Address][]types.Transaction {
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

func (history *HistoryAggregation) calcPrice(t types.Transaction, base types.Address, decimals uint8) (float64, error) {
	if t.Give.Amount.Sign() <= 0 || t.Get.Amount.Sign() <= 0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}

	giveFloat, _ := new(big.Float).SetInt(&t.Give.Amount.Int).Float64()
	getFloat, _ := new(big.Float).SetInt(&t.Get.Amount.Int).Float64()
	decimalsFloat := float64(decimals)

	if t.Get.Token == base {
		getFloat = getFloat / math.Pow(10.0, 18.0)
		giveFloat = giveFloat / math.Pow(10.0, decimalsFloat)
		return (getFloat / giveFloat), nil
	}

	getFloat = getFloat / math.Pow(10.0, decimalsFloat)
	giveFloat = giveFloat / math.Pow(10.0, 18.0)
	return (giveFloat / getFloat), nil
}

func (history *HistoryAggregation) getPrices(transactions []types.Transaction, decimals uint8) ([]float64, []uint) {
	var prices []float64
	var txindex []uint
	for _, tt := range transactions {
		newPrice, _ := history.calcPrice(tt, types.HexToAddress(types.ETH_ADDRESS), decimals)
		prices = append(prices, newPrice)
		txindex = append(txindex, tt.TransactionIndex)
	}

	return prices, txindex
}
