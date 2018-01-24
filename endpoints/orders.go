package endpoints

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

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

	if token == "0x0000000000000000000000000000000000000000" || !common.IsHexAddress(token) {
		// @todo error body
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := getLimit(query.Get("limit"))
	user := getUser(query.Get("user"))

	o := types.Orders{}
	address := types.HexToAddress(token)

	o.Asks = orders.OrderBook.Asks(address, user, limit)
	o.Bids = orders.OrderBook.Bids(address, user, limit)

	json.NewEncoder(rw).Encode(o)
}

func (orders *Orders) GetOrder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	o := orders.OrderBook.GetOrderByHash(params["order"])

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

	//ok, err := orders.BalanceValidator.CheckBalance(o)
	//if err != nil {
	//	log.Printf("checking balance failed: %v", err)
	//	rw.WriteHeader(http.StatusInternalServerError)
	//	// @todo
	//	return
	//}

	//if !ok {
	//	rw.WriteHeadegr(http.StatusBadRequest)
	//	// @todo
	//	return
	//}

	if !common.IsHexAddress(o.Get.Token.String()) || !common.IsHexAddress(o.Give.Token.String()) {
		log.Printf("address is not hex address")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.ToLower(o.Give.Token.String()) == strings.ToLower(o.Get.Token.String()) {
		log.Printf("addresses are identical")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// @todo validate that amounts are not 0

	hash, err := o.OrderHash()
	log.Printf("order hash is: %v", common.ToHex(hash))
	if err != nil {
		log.Printf("hashing order failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		// @todo
		return
	}

	o.Hash = common.ToHex(hash)
	price := calculatePrice(o)

	o.Price = price
	err = orders.OrderBook.InsertOrder(o)
	if err != nil {
		log.Printf("InsertOrder failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func calculatePrice(order types.Order) string {

	get := new(big.Float).SetInt(&order.Get.Amount)
	give := new(big.Float).SetInt(&order.Give.Amount)

	price := new(big.Float)
	if order.Get.Token.IsZero() {
		price = price.Quo(get, give)
	} else {
		price = price.Quo(give, get)
	}

	return price.String()
}

func getLimit(limit string) int {
	if len(limit) != 0 && limit != "0" {

		u, err := strconv.Atoi(limit)
		if err == nil {
			return u
		}
	}

	return 100
}

func getUser(user string) *types.Address {
	if user == "" || !common.IsHexAddress(user) {
		return nil
	}

	addr := types.HexToAddress(user)
	return &addr
}
