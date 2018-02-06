package history

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBName   = "TradeHistory"
	FileName = "History"
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

func (history *MongoHistory) AggregateTransactions(block int64) ([]types.Tick, error) {
	session := history.session.Clone()
	defer session.Close()

	ethAddress := types.HexToAddress("0x0000000000000000000000000000000000000000")
	c := session.DB(DBName).C(FileName)

	match := bson.M{
		"$match": bson.M{"block": block},
	}
	timestampSort := bson.M{
		"$sort": bson.M{"timestamp": 1},
	}
	tokenGroup := bson.M{
		"$project": bson.M{
			"give.token": bson.M{
				"$filter": bson.M{
					"input": "$give.token",
					"as":    "give.tokens",
					"cond": bson.M{"$or": []interface{}{
						bson.M{"$eq": []interface{}{"$$give.tokens", "$get.token"}},
						bson.M{"$eq": []interface{}{"$$give.tokens", "$give.token"}},
						bson.M{"$eq": []interface{}{"$$give.tokens", ethAddress}},
					},
					},
				},
			},
			"get.token": bson.M{
				"$filter": bson.M{
					"input": "$get.token",
					"as":    "$get.tokens",
					"cond": bson.M{"$or": []interface{}{
						bson.M{"$eq": []interface{}{"$$get.tokens", "$give.token"}},
						bson.M{"$eq": []interface{}{"$$get.tokens", "$get.token"}},
						bson.M{"$eq": []interface{}{"$$get.tokens", ethAddress}},
					},
					},
				},
			},
		},
	}

	priceCalc := bson.M{
		"$group": bson.M{
			"_id":       "$block",
			"opentime":  bson.M{"$first": "$timestamp"},
			"closetime": bson.M{"$last": "$timestamp"},
			"getvolume": bson.M{"$sum": bson.M{
				"$cond": []interface{}{bson.M{
					"$eq": []interface{}{ethAddress, "$get.token"},
				},
					0, "$get.amount",
				},
			},
			},
			"givevolume": bson.M{"$sum": bson.M{
				"$cond": []interface{}{bson.M{
					"$eq": []interface{}{ethAddress, "$give.token"},
				},
					0, "$give.amount",
				},
			},
			},
			"price": bson.M{
				"$cond": []interface{}{bson.M{
					"$eq": []interface{}{ethAddress, "$get.token"},
				},
					bson.M{"$divide": []interface{}{"$get.amount", "$give.amount"}},
					bson.M{"$divide": []interface{}{"$give.amount", "$get.amount"}},
				},
			},
		},
	}
	aggregate := bson.M{
		"$group": bson.M{
			"_id":       "$block",
			"opentime":  "$opentime",
			"closetime": "$closetime",
			"volume":    bson.M{"$add": []interface{}{"$givevolume", "$getvolume"}},
			"open":      bson.M{"$first": "$price"},
			"close":     bson.M{"$last": "$price"},
			"high":      bson.M{"$max": "$price"},
			"low":       bson.M{"$min": "$price"},
		},
	}

	pipeline := c.Pipe([]bson.M{match, timestampSort, tokenGroup, priceCalc, aggregate})
	var response []types.Tick
	err := pipeline.All(&response)
	if err != nil {
		return nil, fmt.Errorf("could not query history")
	}

	return response, nil
}
