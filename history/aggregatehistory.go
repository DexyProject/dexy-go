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
	var txindex []float64 //temp


	matchBlock := bson.M{"$match": bson.M{"$transactions.block": block}}
	sortTimestamp := bson.M{"$sort": bson.M{"$transactions.timestamp": -1}}
	groupGetToken := bson.M{
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
	groupGiveToken := bson.M{
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
	groupGiveGetToken := bson.M{
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

	err := c.Pipe([]bson.M{matchBlock, sortTimestamp, groupGetToken, groupGiveToken, groupGiveGetToken}).All(&transactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}
	for _, tt := range transactions {
		pair := getPair(tt.Transactions)
		volume := calcVolume(tt.Transactions)
		prices := getPrices(tt.Transactions)
		openPrices, closePrices := calcOpenClose(txindex)
		high, low := calcHighLow(tt.Transactions, prices)

		ticks = append(ticks, types.Tick{Pair: pair, Block: block, Volume: types.Int{*volume}, Open: openPrices,
		Close: closePrices, High: high, Low: low})
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

func calcHighLow(transactions []types.Transaction, price []float64) (float64, float64) {
	var high, low float64
	for _, p := range price {
		if high > p {
			high = p
		}
	}
	for _, p := range price {
		if low < p {
			low = p
		}
	}

	return high, low
}

func getPrices(transactions []types.Transaction) []float64 {
	var prices []float64
	for _, tt := range transactions {
		newPrice, _ := tt.Get.CalcPrice(tt.Give, types.HexToAddress(types.ETH_ADDRESS))
		prices = append(prices, newPrice)
	}

	return prices
}

func calcOpenClose(txindex []float64) (float64, float64) { //temporary Calculation for open and close until index format is created
	var openPrice, closePrice float64
	openPrice = txindex[len(txindex)-1]
	closePrice = txindex[0]

	return openPrice, closePrice
}

func getPair(transactions []types.Transaction) types.Pair {
	var newPair types.Pair
	for _, tt := range transactions {
		if tt.Give.Token == types.HexToAddress(types.ETH_ADDRESS) {
			newPair = types.Pair{tt.Get.Token, types.HexToAddress(types.ETH_ADDRESS)}
		} else {
			newPair = types.Pair{tt.Give.Token, types.HexToAddress(types.ETH_ADDRESS)}
		}
	}
	return newPair
}
