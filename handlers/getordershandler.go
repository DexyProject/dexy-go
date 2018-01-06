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

	limit := 100
	if query.Get("limit") != "0" {

		u, err := strconv.Atoi(query.Get("limit"))
		if err == nil {
			limit = u
		}
	}

	o := Orders{}

	address := common.HexToAddress(token)
	o.Asks = handler.OrderBook.Asks(address, limit)
	o.Bids = handler.OrderBook.Bids(address, limit)

	json.NewEncoder(rw).Encode(o)
}
