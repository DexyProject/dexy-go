package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/DexyProject/dexy-go/contracts"
	dexyhttp "github.com/DexyProject/dexy-go/http"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

type Orders struct {
	OrderBook        orderbook.OrderBook
	BalanceValidator validators.BalanceValidator
	Vault            *contracts.Vault
}

func (orders *Orders) GetOrderBook(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")
	limit := GetLimit(query.Get("limit"))

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	address := types.HexToAddress(token)

	o := types.Orders{}
	o.Asks = orders.OrderBook.Asks(address, limit)
	o.Bids = orders.OrderBook.Bids(address, limit)

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (orders *Orders) GetOrders(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("user"))

	address := types.HexToAddress(token)

	o := orders.OrderBook.GetOrders(address, user, limit)

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (orders *Orders) GetOrder(rw http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	o := orders.OrderBook.GetOrderByHash(types.NewHash(params["order"]))

	if o == nil {
		http.NotFound(rw, r)
		return nil
	}

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (orders *Orders) CreateOrder(rw http.ResponseWriter, r *http.Request) error {
	var o types.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	defer r.Body.Close()
	if err != nil {
		log.Printf("unmarshalling json failed: %v", err.Error())
		return dexyhttp.NewError("badly formatted order", http.StatusBadRequest)
	}

	approved, err := orders.Vault.IsApproved(nil, o.User.Address, o.Exchange.Address)
	if err != nil {
		log.Printf("checking vault approval failed: %v", err)
		return dexyhttp.NewError("vault approval failed to check", http.StatusInternalServerError)
	}

	if !approved {
		log.Printf("vault is not approved")
		return dexyhttp.NewError("vault is not approved", http.StatusBadRequest)
	}

	ok, err := orders.BalanceValidator.CheckBalance(o)
	if err != nil {
		log.Printf("checking balance failed: %v", err)
		return dexyhttp.NewError("balance check failed", http.StatusInternalServerError)
	}

	if !ok {
		log.Print("insufficient balance to place order")
		return dexyhttp.NewError("insufficient balance to place order", http.StatusBadRequest)
	}

	err = o.Validate()
	if err != nil {
		log.Printf("validating order failed: %v", err)
		return dexyhttp.NewError("validation failed", http.StatusBadRequest)
	}

	price, err := calculatePrice(o)
	if err != nil {
		return dexyhttp.NewError("price error", http.StatusBadRequest)
	}

	o.Price = price
	err = orders.OrderBook.InsertOrder(o)
	if err != nil {
		log.Printf("insert order failed: %v", err)
		return err
	}

	log.Printf("inserted new order %s", o.OrderHash().String())

	rw.WriteHeader(http.StatusCreated)
	return nil
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
