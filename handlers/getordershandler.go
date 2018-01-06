package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type Orders struct {
	Asks []types.Order `json:"asks"`
	Bids []types.Order `json:"bids"`
}

type GetOrdersHandler struct {
	OrderBook orderbook.OrderBook
}

func (handler *GetOrdersHandler) Handle(rw http.ResponseWriter, r *http.Request) {
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

	o := Orders{}
	address := common.HexToAddress(token)

	o.Asks = handler.OrderBook.Asks(address, user, limit)
	o.Bids = handler.OrderBook.Bids(address, user, limit)

	json.NewEncoder(rw).Encode(o)
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
