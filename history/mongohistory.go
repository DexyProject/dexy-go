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

	ethAddress := types.HexToAddress("0x0000000000000000000000000000000000000000")
	c := session.DB(DBName).C(FileName)

	var transactions []types.Transaction


	err := c.Find(bson.M{"block": block}).Sort("-timestamp").All(&transactions)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}

	// group by tokens
	//for i, tt := range transactions {
	//	if tt.Give.Token == tt.Get.Token || tt.Give.Token == ethAddress {
	//		transactions = append(transactions[:i], transactions[i+1:]...)
	//	}
	//}

	// calculate price and volume
	volume := new(big.Int)
	for _, tt := range transactions {
		if tt.Give.Token != ethAddress {
			volume.Add(volume, &tt.Give.Amount.Int)
		}
		if tt.Get.Token != ethAddress {
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}
	var price []*big.Int
	for _, tt := range transactions {
		if tt.Get.Token == ethAddress {

			price = append(price, new(big.Int).Quo(&tt.Get.Amount.Int, &tt.Give.Amount.Int))
		} else {
			price = append(price, new(big.Int).Quo(&tt.Give.Amount.Int, &tt.Get.Amount.Int))
		}
	}

	// find min, max, timestamps and open/close
	var openTime, closeTime int64
	for _, tt := range transactions {
		if openTime < tt.Timestamp {
			openTime = tt.Timestamp
		}
	}
	for _, tt := range transactions {
		if closeTime > tt.Timestamp {
			closeTime = tt.Timestamp
		}
	}
	var open, close *big.Int
	for i, tt := range transactions {
		if tt.Timestamp == closeTime {
			close = price[i]
		}
		if tt.Timestamp == openTime {
			open = price[i]
		}
	}
	var high, low *big.Int
	for _, p := range price {
		if high.Cmp(p) == 1 {
			high = p
		}
	}
	for _, p := range price {
		if low.Cmp(p) == -1 {
			low = p
		}
	}

	var tick = types.Tick{
			Block: block, OpenTime: openTime, CloseTime: closeTime, Volume: types.Int{*volume}, Open: types.Int{*open},
			Close: types.Int{*close} , High: types.Int{*high}, Low: types.Int{*low}}

	return &tick, nil
}
