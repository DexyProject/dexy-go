package history

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
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

func (history *MongoHistory) GetHistory(token common.Address, user *common.Address, limit int) []types.Transaction {
	session := history.session.Clone()
	defer session.Close()

	var transactions []types.Transaction

	c := session.DB(DBName).C(FileName)

	q := bson.M{
		"$or": []bson.M{
			{"give.token": token.String()},
			{"get.token": token.String()},
		},
	}

	if user != nil {
		q["user"] = user.String()
	}

	c.Find(q).Sort("-timestamp").Limit(limit).All(&transactions)

	return transactions
}
