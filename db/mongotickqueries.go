package db

import (
	"gopkg.in/mgo.v2"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

//Tick db queries
type HistoryQueries struct {
	connection string
	session *mgo.Session
}

const (
	DBName = "TickData"
	FileName = "History"
)

func (tq *HistoryQueries) InsertTick(NewTick types.Transaction) error {
	session := tq.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Insert(NewTick)
	if err != nil {
		return fmt.Errorf("could not insert tick data")
	}

	return nil
}

func (tq *HistoryQueries) AggregateTick(block string) ([]bson.M, error) {
	session := tq.session.Clone()
	defer session.Close()

	ethAddress := "0x0000000000000000000000000000000000000000"
	c := session.DB(DBName).C(FileName)

	o1 := bson.M{
		"$match": bson.M{
			"block": block }} //matching block

	o2 := bson.M{
		"$group": bson.M{
			"_id": bson.M {
				"give_token": "$give.token", "get_token": "$get.token"},
			"tickdata": "$$ROOT" }} //@todo verify/test queries

	o3 := bson.M{
		"$match": bson.M{
			"give.token": bson.M{
				"$nin": []interface{}{ ethAddress }},}} // parse out 0x0 from give tokens

	o4 := bson.M{
		"$match": bson.M{
			"get.token": bson.M{
				"$nin": []interface{}{ ethAddress }},}} // parse out 0x0 from get tokens


	pipeline := c.Pipe([]bson.M{o1, o2, o3, o4})
	response := []bson.M{}
	err := pipeline.All(&response)
	if err != nil {
		return nil, fmt.Errorf("could not query history")
	}

	return response, nil
}


