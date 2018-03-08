package ticks

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBName   = "dexy"
	FileName = "ticks"
)

type MongoTicks struct {
	connection string
	session    *mgo.Session
}

func NewMongoTicks(connection string) (*MongoTicks, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to tick database")
	}

	return &MongoTicks{connection: connection, session: session}, nil
}

func (tq *MongoTicks) InsertTicks(ticks []types.Tick) error {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	// @todo insert many
	for i := range ticks {
		err := c.Insert(ticks[i])
		if err != nil {
			return fmt.Errorf("could not insert tick data: %s", err.Error())
		}
	}

	return nil
}

func (tq *MongoTicks) FetchTicks(token types.Address) ([]types.Tick, error) {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	results := make([]types.Tick, 0)

	err := c.Find(bson.M{"pair.quote": token}).Sort("-timestamp").All(&results)
	if err != nil {
		return nil, fmt.Errorf("could not fetch ticks")
	}

	return results, nil
}
