package balances

import (
	"fmt"
	"math/big"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoBalances struct {
	connection string
	session    *mgo.Session
}

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

	c := session.DB(orderbook.DBName).C(orderbook.FileName)

	var result []struct {
		Give struct {
			Amount types.Int `bson:"amount"`
		} `bson:"give"`
	}

	// we solved it like this because mongos $sum function requires values to be numbers, in our case however they are
	// strings.
	err := c.Find(bson.M{"user": user, "give.token": token}).Select(bson.M{"give.amount": 1}).All(&result)
	if err != nil {
		return nil, err
	}

	i := new(big.Int)
	for _, r := range result {
		i = i.Add(i, &r.Give.Amount.Int)
	}

	return &types.Int{Int: *i}, nil
}

func (balances *MongoBalances) Underfunded(user types.Address, token types.Address) (*types.Int, error) {
	return nil, nil
}
