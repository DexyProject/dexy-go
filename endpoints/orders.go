package endpoints

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/DexyProject/dexy-go/contracts"
	dexyhttp "github.com/DexyProject/dexy-go/http"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Orders struct {
	OrderBook        orderbook.OrderBook
	BalanceValidator validators.BalanceValidator
	Vault            *contracts.Vault
}

func (ep *Orders) GetOrderBook(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")
	limit := GetLimit(query.Get("limit"))

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	address := types.HexToAddress(token)

	o := types.Orders{}
	o.Asks = ep.OrderBook.Asks(address, limit)
	o.Bids = ep.OrderBook.Bids(address, limit)

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (ep *Orders) GetOrders(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("maker"))

	address := types.HexToAddress(token)

	o := ep.OrderBook.GetOrders(address, user, limit)

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (ep *Orders) GetOrder(rw http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	o := ep.OrderBook.GetOrderByHash(types.NewHash(params["order"]))

	if o == nil {
		http.NotFound(rw, r)
		return nil
	}

	json.NewEncoder(rw).Encode(o)
	return nil
}

func (ep *Orders) CreateOrder(rw http.ResponseWriter, r *http.Request) error {
	var o types.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	defer r.Body.Close()
	if err != nil {
		log.Debug("unmarshalling json failed", zap.Error(err))
		return dexyhttp.NewError("badly formatted order", http.StatusBadRequest)
	}

	approved, err := ep.Vault.IsApproved(nil, o.Maker.Address, o.Exchange.Address)
	if err != nil {
		log.Error("checking vault approval failed", zap.Error(err))
		return dexyhttp.NewError("vault approval failed to check", http.StatusInternalServerError)
	}

	if !approved {
		log.Debug("vault is not approved")
		return dexyhttp.NewError("vault is not approved", http.StatusBadRequest)
	}

	ok, err := ep.BalanceValidator.CheckBalance(o)
	if err != nil {
		log.Error("checking balance failed", zap.Error(err))
		return dexyhttp.NewError("balance check failed", http.StatusInternalServerError)
	}

	if !ok {
		log.Debug("insufficient balance to place order")
		return dexyhttp.NewError("insufficient balance to place order", http.StatusBadRequest)
	}

	err = o.Validate()
	if err != nil {
		log.Debug("validating order failed", zap.Error(err))
		return dexyhttp.NewError("validation failed", http.StatusBadRequest)
	}

	price, err := calculatePrice(o)
	if err != nil {
		return dexyhttp.NewError("price error", http.StatusBadRequest)
	}

	o.Price = price
	err = ep.OrderBook.InsertOrder(o)
	if err != nil {
		log.Error("insert order failed", zap.Error(err))
		return err
	}

	log.Info("inserted new order", zap.String("hash", o.OrderHash().String()))

	rw.WriteHeader(http.StatusCreated)
	return nil
}

func calculatePrice(order types.Order) (float64, error) {
	if order.Take.Amount.Sign() <= 0 || order.Make.Amount.Sign() <= 0 {
		return 0, fmt.Errorf("can not divide by zero")
	}

	get := new(big.Float).SetInt(&order.Take.Amount.Int)
	give := new(big.Float).SetInt(&order.Make.Amount.Int)

	price := new(big.Float)
	if order.Take.Token.IsZero() {
		calculated, _ := price.Quo(get, give).Float64()
		return calculated, nil
	}
	calculated, _ := price.Quo(give, get).Float64()
	return calculated, nil
}
