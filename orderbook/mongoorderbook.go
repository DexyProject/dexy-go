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
	DBName   = "OrderBook"
	FileName = "Orders"
)

func NewMongoOrderBook(connection string) (*MongoOrderBook, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &MongoOrderBook{connection: connection, session: session}, nil
}

func (ob *MongoOrderBook) InsertOrder(NewOrder types.Order) error {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	if ob.GetOrderByHash(NewOrder.Hash) != nil {
		return fmt.Errorf("order exists in orderbook")
	}

	hash, err := NewOrder.OrderHash()
	if err != nil {
		return fmt.Errorf("could not create order hash")
	}

	if !NewOrder.Signature.Verify(NewOrder.User, hash) {
		return fmt.Errorf("signature could not be verified")
	}

	err = c.Insert(NewOrder)
	if err != nil {
		return err
	}

	return nil
}

func (ob *MongoOrderBook) RemoveOrder(hash string) bool {
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)
	err := c.Remove(bson.M{"hash": hash})
	if err != nil {
		return false
	}

	return true
}

func (ob *MongoOrderBook) Bids(token types.Address, user *types.Address, limit int) []types.Order {
	var orders []types.Order
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"get.token": token}
	if user != nil {
		q["user"] = user
	}

	c.Find(q).Sort("-price").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) Asks(token types.Address, user *types.Address, limit int) []types.Order {
	var orders []types.Order
	session := ob.session.Copy()
	defer session.Close()

	c := session.DB(DBName).C(FileName)

	q := bson.M{"give.token": token}
	if user != nil {
		q["user"] = user
	}

	c.Find(q).Sort("price").Limit(limit).All(&orders)

	return orders
}

func (ob *MongoOrderBook) GetOrderByHash(hash string) *types.Order {
	order := types.Order{}
	session := ob.session.Copy()
	defer session.Close()
	c := session.DB(DBName).C(FileName)

	err := c.Find(bson.M{"hash": hash}).One(&order)
	if err != nil {
		return nil
	}

	return &order
}
