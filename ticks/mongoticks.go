package ticks

import (
	"gopkg.in/mgo.v2"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBName = "TickData"
	FileName = "Ticks"
)

type MongoTicks struct {
	connection string
	session *mgo.Session
}

func NewMongoTicks(connection string) (*MongoTicks, error) {
	session, err := mgo.Dial(connection)
	if err!= nil {
		return nil, fmt.Errorf("could not connect to tick database")
	}

	return &MongoTicks{connection: connection, session: session}, nil
}

func (tq *MongoTicks) InsertTick(NewTick types.Tick) error {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Insert(NewTick)
	if err != nil {
		return fmt.Errorf("could not insert tick data")
	}

	return nil
}

func (tq *MongoTicks) FetchTicks(block int64) ([]types.Tick, error) {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	var results []types.Tick
	err := c.Find(bson.M{"block": block}).All(&results)
	if err!= nil {
		return nil, fmt.Errorf("could not fetch ticks")
	}

	return results, nil
}


