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


func (history *MongoHistory) AggregateTransactions(block int) ([]bson.M, error) {
	session := history.session.Clone()
	defer session.Close()

	ethAddress := "0x0000000000000000000000000000000000000000"
	c := session.DB(DBName).C(FileName)

	o1 := bson.M{
		"$match": bson.M{
			"block": block}} //matching block

	o2 := bson.M{
		"$match": bson.M{
			"give.token": bson.M{
				"$nin": []interface{}{ethAddress}},}} // parse out 0x0 from give tokens

	o3 := bson.M{
		"$match": bson.M{
			"get.token": bson.M{
				"$nin": []interface{}{ethAddress}},}} // parse out 0x0 from get token

	o4 := bson.M{
		"$group": bson.M{
			"_id": "$block",
			"opentime": bson.M{
			},
			"closetime": bson.M{
			},
			"volume": bson.M{
			},
			"open": bson.M{
			},
			"close": bson.M{
			},
			"high": bson.M{
				"$cond": []interface{}{bson.M{
					"$gte": []interface{}{bson.M{
						"$max": "$give.amount"}, bson.M{"$max": "$get.amount"}}},
					bson.M{"$max": "$give.amount"},
					bson.M{"$max": "get.amount"}}},
			"low": bson.M{
				"$cond": []interface{}{bson.M{
					"$gte": []interface{}{bson.M{
						"$min": "$give.amount"}, bson.M{"$min": "$get.amount"}}},
					bson.M{"$min": "$give.amount"},
					bson.M{"$max": "$get.amount"}}}}}


	pipeline := c.Pipe([]bson.M{o1, o2, o3, o4})
	response := []bson.M{}
	err := pipeline.All(&response)
	if err != nil {
		return nil, fmt.Errorf("could not query history")
	}

	return response, nil
}
