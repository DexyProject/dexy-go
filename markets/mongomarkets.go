package markets

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoMarkets struct {
	session *mgo.Session
}

const (
	DBName   = "dexy"
	FileName = "markets"
)

func NewMongoMarkets(connection string) (*MongoMarkets, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to tick database")
	}

	return &MongoMarkets{session: session}, nil
}

func (m *MongoMarkets) InsertMarkets(markets []types.Market) error {
	session := m.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	// @todo insert many
	for i := range markets {
		_, err := c.Upsert(bson.M{"_id": markets[i].Token}, markets[i])
		if err != nil {
			return fmt.Errorf("could not insert market data: %s", err.Error())
		}
	}

	return nil
}

func (m *MongoMarkets) GetMarkets(tokens []types.Address) ([]types.Market, error) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	markets := make([]types.Market, 0)

	err := c.Find(bson.M{"_id": bson.M{"$in": tokens}}).All(&markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}
