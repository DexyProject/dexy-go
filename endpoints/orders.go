package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

type Orders struct {
	OrderBook        orderbook.OrderBook
	BalanceValidator validators.BalanceValidator
}

func (orders *Orders) GetOrderBook(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	token := query.Get("token")
	limit := GetLimit(query.Get("limit"))

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		returnError(rw, "invalid token", http.StatusBadRequest)
		return
	}

	address := types.HexToAddress(token)

	o := types.Orders{}
	o.Asks = orders.OrderBook.Asks(address, limit)
	o.Bids = orders.OrderBook.Bids(address, limit)

	json.NewEncoder(rw).Encode(o)
}

func (orders *Orders) GetOrders(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		returnError(rw, "invalid token", http.StatusBadRequest)
		return
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("user"))

	address := types.HexToAddress(token)

	o := orders.OrderBook.GetOrders(address, user, limit)

	json.NewEncoder(rw).Encode(o)
}

func (orders *Orders) GetOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	o := orders.OrderBook.GetOrderByHash(types.NewHash(params["order"]))

	if o == nil {
		http.NotFound(rw, r)
		return
	}

	json.NewEncoder(rw).Encode(o)
}

func (orders *Orders) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var o types.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	defer r.Body.Close()
	if err != nil {
		log.Printf("unmarshalling json failed: %v", err.Error())
		returnError(rw, "badly formatted order", http.StatusBadRequest)
		return
	}

	ok, err := orders.BalanceValidator.CheckBalance(o)
	if err != nil {
		log.Printf("checking balance failed: %v", err)
		returnError(rw, "balance check failed", http.StatusInternalServerError)
		return
	}

	if !ok {
		log.Print("insufficient balance to place order")
		returnError(rw, "insufficient balance to place order", http.StatusBadRequest)
		return
	}

	err = o.Validate()
	if err != nil {
		log.Printf("validating order failed: %v", err)
		returnError(rw, "validation failed", http.StatusBadRequest)
		return
	}

	price, err := calculatePrice(o)
	if err != nil {
		returnError(rw, "price error", http.StatusBadRequest)
		return
	}

	o.Price = price
	err = orders.OrderBook.InsertOrder(o)
	if err != nil {
		log.Printf("InsertOrder failed: %v", err)
		returnError(rw, "internal error", http.StatusInternalServerError)
		return
	}

	log.Printf("inserted new order %s", o.OrderHash().String())

	rw.WriteHeader(http.StatusCreated)
}

func calculatePrice(order types.Order) (string, error) {

	if order.Get.Amount.Sign() <= 0 || order.Give.Amount.Sign() <= 0 {
		return "", fmt.Errorf("can not divide by zero")
	}

	get := new(big.Float).SetInt(&order.Get.Amount.Int)
	give := new(big.Float).SetInt(&order.Give.Amount.Int)

	price := new(big.Float)
	if order.Get.Token.IsZero() {
		return price.Quo(get, give).Text('f', 8), nil
	}

	return price.Quo(give, get).Text('f', 8), nil
}
