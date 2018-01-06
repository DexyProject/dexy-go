package orderbook

import (
	"github.com/DexyProject/dexy-go/types"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/ethereum/go-ethereum/common"
)

type MongoOrderBook struct {
	connection string
}

const (
	DBName = "OrderBook"
	FileName = "Orders"
)

func NewMongoDataProvider(connection string) (mgo.Session, error) {
	session, err := mgo.Dial(connection)
	if err !=  nil {
		return *session, fmt.Errorf("connection to db could not be established")
	}
	return *session, err

}
func (ob *MongoOrderBook) InsertOrder(NewOrder types.Order) error {
	// Connect to Mongo session
	session, _ := NewMongoDataProvider(ob.connection)
	c := session.DB(DBName).C(FileName)

	if ob.GetOrderByHash(NewOrder.Hash) != nil { // Check if Hash exists
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
		return fmt.Errorf("order could not be added to database")
	}
	defer session.Close()
	return nil
}

func (ob *MongoOrderBook) RemoveOrder(hash string) bool {
	session, _ := NewMongoDataProvider(ob.connection)
	c := session.DB(DBName).C(FileName)
	err := c.Remove(bson.M{"hash": hash})
	if err != nil {
		return false
	}
	defer session.Close()
	return true
}


func (ob *MongoOrderBook) Bids(token common.Address, limit int, user *common.Address) ([]types.Order) {

	var orders []types.Order
	session, _ := NewMongoDataProvider(ob.connection)
	c := session.DB(DBName).C(FileName)
	if user != nil {
		c.Find(bson.M{"user":user}).Sort("-price").Limit(limit).All(&orders)
	} else {
		c.Find(bson.M{"token": token}).Sort("-price").Limit(limit).All(&orders)
	}
	defer session.Close()
	return orders
}

func (ob *MongoOrderBook) Asks(token common.Address, limit int, user *common.Address) ([]types.Order) {

	var orders []types.Order
	session, _ := NewMongoDataProvider(ob.connection)
	c := session.DB(DBName).C(FileName)
	if user != nil {
		c.Find(bson.M{"user":user}).Sort("price").Limit(limit).All(&orders)
	} else {
		c.Find(bson.M{"token": token}).Sort("price").Limit(limit).All(&orders)
	}
	defer session.Close()
	return orders
}

func (ob *MongoOrderBook) GetOrderByHash(hash string) (*types.Order) {
	order := types.Order{}
	session, _ := NewMongoDataProvider(ob.connection)
	c := session.DB(DBName).C(FileName)
	err := c.Find(bson.M{"hash":hash}).One(&order)
	if err != nil {
		return nil
	}
	defer session.Close()
	return &order
}
