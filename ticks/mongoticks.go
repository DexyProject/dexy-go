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
	session *mgo.Session
}

func NewMongoTicks(connection string) (*MongoTicks, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to tick database")
	}

	return &MongoTicks{session: session}, nil
}

func (t *MongoTicks) InsertTicks(ticks []types.Tick) error {
	session := t.session.Clone()
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

func (t *MongoTicks) FetchTicks(token types.Address) ([]types.Tick, error) {
	session := t.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	results := make([]types.Tick, 0)

	err := c.Find(bson.M{"pair.quote": token}).Sort("-timestamp").All(&results)
	if err != nil {
		return nil, fmt.Errorf("could not fetch ticks")
	}

	return results, nil
}

//func (t *MongoTicks) FetchLatestTickForTokens(tokens []types.Address) (map[types.Address]types.Tick, error) {
//	session := t.session.Clone()
//	defer session.Close()
//
//	c := session.DB(DBName).C(FileName)
//	results := make(map[types.Address]types.Tick, 0)
//
//	pipe := []bson.M{
//		{"$match": bson.M{"pair.quote": bson.M{"$in": tokens}}},
//		{"$sort": bson.M{"timestamp": -1}},
//		{"$group": bson.M{"_id": "$pair.quote", "ticks": bson.M{"$push": "$$ROOT"}}},
//		{"$replaceRoot": bson.M{"newRoot": bson.M{"$arrayElemAt": []interface{}{"$ticks", 0}}}},
//	}
//
//	ticks := make([]types.Tick, 0)
//	err := c.Pipe(pipe).All(&ticks)
//	if err != nil {
//		return results, err
//	}
//
//	for _, tick := range ticks {
//		results[tick.Pair.Quote] = tick
//	}
//
//	return results, nil
//}

func (t *MongoTicks) FetchAggregateVolumeForTokens(tokens []types.Address) (map[types.Address]types.Int, error) {
	// @todo
}


func (t *MongoTicks) FetchLatestCloseForTokens(tokens []types.Address) (map[types.Address]types.Int, error) {
	// @todo
}
