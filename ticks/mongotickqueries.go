package ticks

import (
	"gopkg.in/mgo.v2"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
)

//Tick ticks queries
type MongoTickQueries struct {
	connection string
	session *mgo.Session
}

const (
	DBName = "TickData"
	FileName = "Ticks"
)

func (tq *MongoTickQueries) InsertTick(NewTick types.Transaction) error {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Insert(NewTick)
	if err != nil {
		return fmt.Errorf("could not insert tick data")
	}

	return nil
}

func (tq *MongoTickQueries) FetchTicks() {
	//todo
}


