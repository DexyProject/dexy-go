package orderbook

import (
	"fmt"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoOrderBook struct {
	connection string
	session    *mgo.Session
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

	return &MongoOrderBook{connection: connection, session: session}, nil
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
			{"give.token": token},
			{"get.token": token},
		},
	}

	if user != nil {
		q["user"] = user
	}

	orders := make([]types.Order, 0)
	c.Find(q).Sort("-expires").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) Bids(token types.Address, limit int) []types.Order {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"get.token": token}

	orders := make([]types.Order, 0)
	c.Find(q).Sort("-price").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) Asks(token types.Address, limit int) []types.Order {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"give.token": token}

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

// @todo this is ugly, find a cleaner way at a later stage, possibly move out of OB
func (ob *MongoOrderBook) GetMarkets(tokens []types.Address) (map[types.Address]*types.Market, error) {
	m := make(map[types.Address]*types.Market)

	markets, err := ob.getMarkets(tokens)
	if err != nil {
		return m, err
	}

	for _, market := range markets {
		ask := market["ask"].(bson.M)
		bid := market["bid"].(bson.M)

		m[types.HexToAddress(market["token"].(string))] = &types.Market{
			Ask: types.PairAmount{Quote: ask["quote"].(string), Base: ask["base"].(string)},
			Bid: types.PairAmount{Base: bid["base"].(string), Quote: bid["quote"].(string)},
		}
	}

	return m, nil
}

func (ob *MongoOrderBook) getMarkets(tokens []types.Address) ([]bson.M, error) {
	return ob.executeAggregation(
		[]bson.M{
			{
				"$match": bson.M{
					"$or": []bson.M{
						{"give.token": bson.M{"$in": tokens}},
						{"get.token": bson.M{"$in": tokens}},
					},
				},
			},
			{
				"$group": bson.M{
					"_id": bson.M{
						"$cond": bson.M{
							"if":   bson.M{"$eq": []string{"$give.token", types.ETH_ADDRESS}},
							"then": "$get.token",
							"else": "$give.token",
						},
					},
					"bids": bson.M{
						"$push": bson.M{
							"$cond": bson.M{
								"if":   bson.M{"$eq": []string{"$give.token", types.ETH_ADDRESS}},
								"then": bson.M{"base": "$give.amount", "quote": "$get.amount", "price": "$price"},
								"else": bson.M{"base": "0", "quote": "0", "price": "0"},
							},
						},
					},
					"asks": bson.M{
						"$push": bson.M{
							"$cond": bson.M{
								"if":   bson.M{"$eq": []string{"$get.token", types.ETH_ADDRESS}},
								"then": bson.M{"base": "$get.amount", "quote": "$give.amount", "price": "$price"},
								"else": bson.M{"base": "0", "quote": "0", "price": "0"},
							},
						},
					},
				},
			},
			{
				"$project": bson.M{
					"_id":   0,
					"token": "$_id",
					"ask":   bson.M{"$arrayElemAt": []interface{}{"$asks", 0}},
					"bid":   bson.M{"$arrayElemAt": []interface{}{"$bids", -1}},
				},
			},
		},
	)
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
