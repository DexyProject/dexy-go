package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	address := common.HexToAddress(token)

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

	decoder := json.NewDecoder(r.Body)
	var o types.Order
	err := decoder.Decode(o)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		// @todo
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

	hash, err := o.OrderHash()
	if err != nil {
		log.Printf("hashing order failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		// @todo
		return
	}

	o.Hash = common.ToHex(hash)
	price, err := calculatePrice(o)
	if err != nil {
		log.Printf("price calculation failed: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		// @todo
		return
	}

	o.Price = price
	err = orders.OrderBook.InsertOrder(o)
	if err != nil {
		return
	}
	// @todo response
}

func calculatePrice(order types.Order) (string, error) {

	get, err := strconv.ParseFloat(order.Get.Amount, 64)
	if err != nil {
		return "", err
	}

	give, err := strconv.ParseFloat(order.Give.Amount, 64)
	if err != nil {
		return "", err
	}

	var price float64
	if order.Get.Token.String() == "0x0000000000000000000000000000000000000000" {
		price = get / give
	} else {
		price = give / get
	}

	return strconv.FormatFloat(price, 'f', -1, 64), nil
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

func getUser(user string) *common.Address {
	if user == "" || !common.IsHexAddress(user) {
		return nil
	}

	addr := common.HexToAddress(user)
	return &addr
}
