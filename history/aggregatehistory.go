package history

import (
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"math/big"
)

type Transactions struct {
	Transactions   []types.Transaction `json:"transactions" bson:"transactions"`
}

func (history *MongoHistory) AggregateTransactions(block int64, transactions []Transactions) ([]types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	var ticks []types.Tick

	matchBlock := bson.M{"$match": bson.M{"$transactions.block": block}}

	err := c.Pipe([]bson.M{matchBlock}).All(&transactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}
	for _, tt := range transactions {
		groupTokens(tt.Transactions)
		pair := getPair(tt.Transactions)
		volume := calcVolume(tt.Transactions)
		prices := getPrices(tt.Transactions)
		openIndex, closeIndex := calcOpenCloseIndex(tt.Transactions)
		openPrice, closePrice := calcOpenClosePrice(prices, openIndex, closeIndex)
		high, low := calcHighLow(prices)

		ticks = append(ticks, types.Tick{Pair: pair, Block: block, Volume: types.Int{*volume}, Open: openPrice,
		Close: closePrice, High: high, Low: low})
	}

	return ticks, nil
}

func calcVolume(transactions []types.Transaction) *big.Int {
	volume := new(big.Int)
	for _, tt := range transactions {
		if tt.Give.Token != types.HexToAddress(types.ETH_ADDRESS) {
			volume.Add(volume, &tt.Give.Amount.Int)
		}
		if tt.Get.Token != types.HexToAddress(types.ETH_ADDRESS) {
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}

	return volume
}

func calcHighLow(prices []types.Price) (float64, float64) {
	high, low := prices[0].Price, prices[0].Price
	for _, p := range prices {
		if high < p.Price {
			high = p.Price
		}
		if low > p.Price {
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
		if openIndex > tt.TransactionIndex {
			openIndex = tt.TransactionIndex
		}
		if closeIndex < tt.TransactionIndex {
			closeIndex = tt.TransactionIndex}
	}

	return openIndex, closeIndex
}

func calcOpenClosePrice(prices []types.Price, OpenIndex, CloseIndex uint) (float64, float64) {
	var openPrice, closePrice float64
	for _, tt := range prices {
		if tt.TransactionIndex == OpenIndex {
			openPrice = tt.Price
		} else if tt.TransactionIndex == CloseIndex {
			closePrice = tt.Price
		}
	}

	return openPrice, closePrice
}

func getPair(transactions []types.Transaction) types.Pair {
	var newPair types.Pair
	if transactions[1].Give.Token == types.HexToAddress(types.ETH_ADDRESS) {
		newPair = types.Pair{transactions[1].Get.Token, types.HexToAddress(types.ETH_ADDRESS)}
	}
	newPair = types.Pair{types.HexToAddress(types.ETH_ADDRESS), transactions[1].Get.Token}
	return newPair
}

func groupTokens(transactions []types.Transaction) {
	// x == x, x!=base, x==y, and x!= base, y=y and y!=base
	for _, t := range transactions {
		for i, x := range transactions {
			if t.Give.Token == x.Give.Token && t.Give.Token != types.HexToAddress(types.ETH_ADDRESS) {
				copy(transactions[i:], transactions[i+1:])
				transactions[len(transactions)-1] = types.Transaction{}
				transactions = transactions[:len(transactions)-1]
			}
			if t.Give.Token == x.Get.Token && t.Give.Token != types.HexToAddress(types.ETH_ADDRESS) {
				copy(transactions[i:], transactions[i+1:])
				transactions[len(transactions)-1] = types.Transaction{}
				transactions = transactions[:len(transactions)-1]
			}
			if t.Get.Token == x.Get.Token && t.Get.Token != types.HexToAddress(types.ETH_ADDRESS) {
				copy(transactions[i:], transactions[i+1:])
				transactions[len(transactions)-1] = types.Transaction{}
				transactions = transactions[:len(transactions)-1]
			}
		}
	}
}