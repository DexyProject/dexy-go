package balances

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoBalances struct {
	connection string
	session    *mgo.Session
}

const (
	DBName   = "OrderBook"
	FileName = "Orders"
)

func NewMongoBalances(connection string) (*MongoBalances, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &MongoBalances{connection: connection, session: session}, nil
}

func (balances *MongoBalances) OnOrders(user types.Address, token types.Address) (*types.Int, error) {
	session := balances.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	match := bson.M{"$match": bson.M{"user": user, "give.token": token}}

	price := bson.M{
		"$group": bson.M{
			"_id": nil,
			"amount": bson.M{"$sum": "$give.amount"},
		},
	}

	pipe := c.Pipe([]bson.M{match, price})

	var result struct {
		Amount types.Int `bson:"amount"`
	}

	err := pipe.One(&result)
	if err != nil {
		return nil, err
	}

	return &result.Amount, nil
}
