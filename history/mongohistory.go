package history

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/big"
)

const (
	DBName   = "tradehistory"
	FileName = "history"

)

type MongoHistory struct {
	connection string
	session    *mgo.Session
}

func NewMongoHistory(connection string) (*MongoHistory, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &MongoHistory{connection: connection, session: session}, nil
}

func calcOpenTime(transactions []types.Transaction) int64 {
	var openTime int64
	for _, tt := range transactions {
		if openTime < tt.Timestamp {
			openTime = tt.Timestamp
		}
	}

	return openTime
}

func calcCloseTime(transactions []types.Transaction) int64 {
	var closeTime int64

	for _, tt := range transactions {
		if closeTime > tt.Timestamp {
			closeTime = tt.Timestamp
		}
	}

	return closeTime
}

func calcVolume(transactions []types.Transaction) *big.Int {
	ethAddress := types.HexToAddress("0x0000000000000000000000000000000000000000")
	volume := new(big.Int)

	for _, tt := range transactions {
		if tt.Give.Token != ethAddress {
			volume.Add(volume, &tt.Give.Amount.Int)
		}
		if tt.Get.Token != ethAddress {
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}

	return volume
}

func calcPrice(transactions []types.Transaction) []float64 {
	ethAddress := types.HexToAddress("0x0000000000000000000000000000000000000000")
	var price []float64

	for _, tt := range transactions {
		if tt.Get.Token == ethAddress {
			newPrice := new(big.Int).Quo(&tt.Get.Amount.Int, &tt.Give.Amount.Int)
			newFloat, _ := new(big.Float).SetInt(newPrice).Float64()
			price = append(price, newFloat)
		} else {
			newPrice := new(big.Int).Quo(&tt.Give.Amount.Int, &tt.Get.Amount.Int)
			newFloat, _ := new(big.Float).SetInt(newPrice).Float64()
			price = append(price, newFloat)
		}
	}

	return price
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

func calcOpenClose(transactions []types.Transaction, price []float64, openTime, closeTime int64) (float64, float64) {
	var openPrice, closePrice float64
	for i, tt := range transactions {
		if tt.Timestamp == closeTime {
			closePrice = price[i]
		}
		if tt.Timestamp == openTime {
			openPrice = price[i]
		}
	}

	return openPrice, closePrice
}

func (history *MongoHistory) GetHistory(token types.Address, user *types.Address, limit int) []types.Transaction {
	session := history.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	var transactions []types.Transaction

	q := bson.M{
		"$or": []bson.M{
			{"give.token": token},
			{"get.token": token},
		},
	}

	if user != nil {
		q["user"] = user
	}

	c.Find(q).Sort("-timestamp").Limit(limit).All(&transactions)

	return transactions
}

func (history *MongoHistory) AggregateTransactions(block int64) (*types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	var transactions []types.Transaction

	err := c.Find(bson.M{"block": block}).Sort("-timestamp").All(&transactions)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}
	openTime := calcOpenTime(transactions)
	closeTime := calcCloseTime(transactions)
	volume := calcVolume(transactions)
	price := calcPrice(transactions)
	openPrice, closePrice := calcOpenClose(transactions, price, openTime, closeTime)
	high, low := calcHighLow(transactions, price)

	var tick = types.Tick{
			Block: block, OpenTime: openTime, CloseTime: closeTime, Volume: types.Int{*volume}, Open: openPrice,
			Close: closePrice , High: high, Low: low}

	return &tick, nil
}
