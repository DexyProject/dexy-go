package markets

import (
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

func (m *MongoMarkets) InsertMarkets(markets []types.Market) error {
	panic("implement me")
}

func (m *MongoMarkets) GetMarkets(tokens []types.Address) ([]types.Market, error) {
	session := m.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	markets := make([]types.Market, 0)

	err := c.Find(bson.M{"token": bson.M{"$in": tokens}}).All(markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}
