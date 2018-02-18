package history

import (
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"math/big"
)

func (history *MongoHistory) AggregateTransactions(block int64) ([]types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	var transactions []struct {
		Transactions []types.Transaction `json:"transactions" bson:"transactions"`
	}

	var ticks []types.Tick

	matchBlock := bson.M{"$match": bson.M{"$transactions.block": block}}
	sortTimestamp := bson.M{"$sort": bson.M{"$transactions.timestamp": -1}}
	groupGetTokens := bson.M{
		"$group": bson.M{
			"$filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.get.token", "$$tt.get.token"}},
					bson.M{"$ne": []interface{}{"$$tt.get.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}
	groupGiveTokens := bson.M{
		"$group": bson.M{
			"$filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.give.token", "$$tt.give.token"}},
					bson.M{"$ne": []interface{}{"$$tt.give.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}
	groupTokens := bson.M{
		"$group": bson.M{
			"$filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.give.token", "$$tt.get.token"}},
					bson.M{"$ne": []interface{}{"$$tt.get.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}

	err := c.Pipe([]bson.M{matchBlock, sortTimestamp, groupGetTokens, groupGiveTokens, groupTokens}).All(&transactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}
	for _, tt := range transactions {
		pair := GetPair(tt.Transactions)
		volume := CalcVolume(tt.Transactions)
		prices := GetPrices(tt.Transactions)
		openIndex, closeIndex := CalcOpenCloseIndex(tt.Transactions)
		openPrice, closePrice := CalcOpenClosePrice(prices, openIndex, closeIndex)
		high, low := CalcHighLow(prices)

		ticks = append(ticks, types.Tick{Pair: pair, Block: block, Volume: types.Int{*volume}, Open: openPrice,
		Close: closePrice, High: high, Low: low})
	}

	return ticks, nil
}

func CalcVolume(transactions []types.Transaction) *big.Int {
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

func CalcHighLow(prices []types.Price) (float64, float64) {
	high, low := prices[0].Price, prices[0].Price
	for _, p := range prices {
		if high > p.Price {
			high = p.Price
		}
		if low < p.Price {
			low = p.Price
		}
	}

	return high, low
}

func GetPrices(transactions []types.Transaction) []types.Price {
	var prices []types.Price
	for _, tt := range transactions {
		newPrice, _ := tt.Get.CalcPrice(tt.Give, types.HexToAddress(types.ETH_ADDRESS))
		priceStruct := types.Price{tt.TransactionIndex, newPrice}
		prices = append(prices, priceStruct)
	}

	return prices
}

func CalcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
	openIndex, closeIndex := transactions[0].TransactionIndex, transactions[0].TransactionIndex
	for _, tt := range transactions {
		if openIndex < tt.TransactionIndex {
			openIndex = tt.TransactionIndex
		}
		if closeIndex > tt.TransactionIndex {
			closeIndex = tt.TransactionIndex}
	}


	return openIndex, closeIndex
}

func CalcOpenClosePrice(prices []types.Price, OpenIndex, CloseIndex uint) (float64, float64) { //temporary Calculation for open and close until index format is created
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

func GetPair(transactions []types.Transaction) types.Pair {
	var newPair types.Pair
	if transactions[1].Give.Token == types.HexToAddress(types.ETH_ADDRESS) {
		newPair = types.Pair{transactions[1].Get.Token, types.HexToAddress(types.ETH_ADDRESS)}
	}
	newPair = types.Pair{types.HexToAddress(types.ETH_ADDRESS), transactions[1].Get.Token}
	return newPair
}
