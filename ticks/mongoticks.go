package ticks

import (
	"fmt"
	"math/big"
	"time"

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

func (t *MongoTicks) FetchAggregateVolumeForTokens(tokens []types.Address) (map[types.Address]types.Int, error) {
	results := make(map[types.Address]types.Int, 0)

	data, err := t.executeAggregation(
		[]bson.M{
			{
				"$match": bson.M{
					"pair.quote": bson.M{"$in": tokens},
					"timestamp":  bson.M{"$gt": types.Timestamp{Time: time.Now().Add(-24 * time.Hour)}},
				},
			},
			{"$sort": bson.M{"timestamp": -1}},
			{"$group": bson.M{"_id": "$pair.quote", "volume": bson.M{"$push": "$volume"}}},
		},
	)

	if err != nil {
		return results, err
	}

	for _, tick := range data {
		vol, err := calculateVolume(tick["volume"].([]interface{}))
		if err != nil {
			return results, err
		}

		results[types.HexToAddress(tick["_id"].(string))] = *vol
	}

	return results, nil
}

func (t *MongoTicks) FetchLatestCloseForTokens(tokens []types.Address) (map[types.Address]float64, error) {
	results := make(map[types.Address]float64, 0)

	data, err := t.executeAggregation(
		[]bson.M{
			{"$match": bson.M{"pair.quote": bson.M{"$in": tokens}}},
			{"$sort": bson.M{"timestamp": -1}},
			{"$group": bson.M{"_id": "$pair.quote", "close": bson.M{"$push": "$close"}}},
			{"$project": bson.M{"token": "$_id", "close": bson.M{"$arrayElemAt": []interface{}{"$close", 0}}}},
		},
	)

	if err != nil {
		return results, err
	}

	for _, tick := range data {
		results[types.HexToAddress(tick["token"].(string))] = tick["close"].(float64)
	}

	return results, nil
}

func (t *MongoTicks) executeAggregation(pipeline interface{}) ([]bson.M, error) {
	session := t.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	pipe := c.Pipe(pipeline)

	var result []bson.M
	err := pipe.All(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func calculateVolume(vols []interface{}) (*types.Int, error) {
	volume := new(big.Int)

	for _, vol := range vols {
		v, ok := new(big.Int).SetString(vol.(string), 10)
		if !ok {
			return nil, fmt.Errorf("could not create volume int for %s", vol.(string))
		}

		volume = volume.Add(volume, v)
	}

	return &types.Int{Int: *volume}, nil
}
