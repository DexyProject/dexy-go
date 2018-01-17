package db

import (
	"gopkg.in/mgo.v2"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//Tick db queries
type TickQueries struct {
	connection string
	session *mgo.Session
}

const (
	DBName = "TickData"
	FileName = "History"
)

func (tq *TickQueries) TickInsert(NewTick types.Tick) error {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Insert(NewTick)
	if err != nil {
		return fmt.Errorf("could not insert tick data")
	}

	return nil
}

func (tq *TickQueries) TickAggregate(block string) ([]bson.M, error) { //Aggregate query data
	session := tq.session.Clone()
	defer session.Close()

	timeNow := time.Now().Unix() // already int64
	tickRate := 3 // seconds
	c := session.DB(DBName).C(FileName)

	o1 := bson.M{
		"$match": bson.M{
			"block": block }} //matching block

	o2 := bson.M{
		"$group": bson.M{
			"_id": "$give.token" }} //aggregate by token

	o3 := bson.M{
		"$project": bson.M{
			"_id": 1, "tx": 1, "hash": 1, "taker": 1, "maker": 1, "give": 1, "get": 1, "timeDiff": bson.M{
				"$subtract": []interface{}{timeNow, "$opentime"}}, "closetime": 1 }}

	o4 := bson.M{
		"$match": bson.M{
			"timeDiff": bson.M{
				"$lte": tickRate }}}

	o5 := bson.M {} //todo //filter 0x0

	pipeline := c.Pipe([]bson.M{o1, o2, o3, o4, o5})
	response := []bson.M{}
	err := pipeline.All(&response)
	if err != nil {
		return nil, fmt.Errorf("could not query history")
	}

	return response, nil
}


