package orderbook

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoOrderBook struct {
	session *mgo.Session
}

const (
	DBName   = "dexy"
	FileName = "orders"
)

func NewMongoOrderBook(connection string) (*MongoOrderBook, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &MongoOrderBook{session: session}, nil
}

func (ob *MongoOrderBook) InsertOrder(order types.Order) error {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	hash := order.OrderHash()
	if ob.GetOrderByHash(hash) != nil {
		return fmt.Errorf("order exists in orderbook")
	}

	if !order.Signature.Verify(order.User, hash) {
		return fmt.Errorf("signature could not be verified (hash %s)", hash.String())
	}

	order.Filled = types.NewInt(0)

	err := c.Insert(order)
	if err != nil {
		return err
	}

	return nil
}

func (ob *MongoOrderBook) RemoveOrder(hash types.Hash) bool {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Remove(bson.M{"_id": hash})
	if err != nil {
		return false
	}

	return true
}

func (ob *MongoOrderBook) GetOrders(token types.Address, user *types.Address, limit int) []types.Order {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{
		"$or": []bson.M{
			{"make.token": token},
			{"take.token": token},
		},
	}

	if user != nil {
		q["maker"] = user
	}

	orders := make([]types.Order, 0)
	c.Find(q).Sort("-expires").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) Bids(token types.Address, limit int) []types.Order {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"take.token": token}

	orders := make([]types.Order, 0)
	c.Find(q).Sort("-price").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) Asks(token types.Address, limit int) []types.Order {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"make.token": token}

	orders := make([]types.Order, 0)
	c.Find(q).Sort("price").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) UpdateOrderFilledAmount(hash types.Hash, amount types.Int) error {
	session := ob.session.Clone()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	err := c.Update(bson.M{"_id": hash}, bson.M{"$set": bson.M{"filled": amount}})
	if err != nil {
		return err
	}

	return nil
}

func (ob *MongoOrderBook) GetOrderByHash(hash types.Hash) *types.Order {
	order := types.Order{}
	session := ob.session.Copy()
	defer session.Close()
	c := session.DB(DBName).C(FileName)

	err := c.Find(bson.M{"_id": hash}).One(&order)
	if err != nil {
		return nil
	}

	return &order
}

func (ob *MongoOrderBook) GetLowestAsks(tokens []types.Address) (types.Prices, error) {
	return ob.getPricesForOrder(
		bson.M{"make.token": bson.M{"$in": tokens}},
		1,
		bson.M{
			"_id":  "$make.token",
			"data": bson.M{"$push": bson.M{"base": "$take.amount", "quote": "$make.amount"}},
		},
	)
}

func (ob *MongoOrderBook) GetHighestBids(tokens []types.Address) (types.Prices, error) {
	return ob.getPricesForOrder(
		bson.M{"take.token": bson.M{"$in": tokens}},
		-1,
		bson.M{
			"_id":  "$take.token",
			"data": bson.M{"$push": bson.M{"base": "$make.amount", "quote": "$take.amount"}},
		},
	)
}

func (ob *MongoOrderBook) getPricesForOrder(match bson.M, sort int, group bson.M) (types.Prices, error) {
	p := make(types.Prices)

	result, err := ob.executeAggregation(
		[]bson.M{
			{"$match": match},
			{"$sort": bson.M{"price": sort}},
			{"$group": group},
			{"$project": bson.M{"token": "$_id", "data": bson.M{"$arrayElemAt": []interface{}{"$data", 0}}}},
			{"$project": bson.M{"token": "$_id", "base": "$data.base", "quote": "$data.quote"}},
		},
	)

	if err != nil {
		return p, err
	}

	for _, data := range result {

		p[types.HexToAddress(data["token"].(string))] = types.Price{
			Base:  data["base"].(string),
			Quote: data["quote"].(string),
		}
	}

	return p, nil
}

func (ob *MongoOrderBook) executeAggregation(pipeline interface{}) ([]bson.M, error) {
	session := ob.session.Copy()
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
