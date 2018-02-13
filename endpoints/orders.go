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

func (orders *Orders) GetOrders(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		// @todo error body
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("user"))

	o := types.Orders{}
	address := types.HexToAddress(token)

	o.Asks = orders.OrderBook.Asks(address, user, limit)
	o.Bids = orders.OrderBook.Bids(address, user, limit)

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
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ok, err := orders.BalanceValidator.CheckBalance(o)
	if err != nil {
		log.Printf("checking balance failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		// @todo
		return
	}

	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		log.Print("insufficient balance to place order")
		// @todo
		return
	}

	err = o.Validate()
	if err != nil {
		log.Printf("validating order failed: %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	price, err := calculatePrice(o)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	o.Price = price
	err = orders.OrderBook.InsertOrder(o)
	if err != nil {
		log.Printf("InsertOrder failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

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
