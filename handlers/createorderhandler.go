package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type CreateOrderHandler struct {
	OrderBook orderbook.OrderBook
}

func (handler *CreateOrderHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var o types.Order
	err := decoder.Decode(o)
	if err != nil {
		return
		// @todo
	}

	hash, err := o.OrderHash()
	if err != nil {
		// @todo
		return
	}

	o.Hash = common.ToHex(hash)

	err = handler.OrderBook.InsertOrder(o)
	if err != nil {
		return
	}

	// @todo response
}
