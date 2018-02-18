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

func (history *MongoHistory) InsertTransaction(transaction types.Transaction) error {
	session := history.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	err := c.Insert(transaction)
	if err != nil {
		return err
	}

	return nil
}
