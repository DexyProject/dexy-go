package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/decanus/dexy-go/orderbook"
	"github.com/gorilla/mux"
)

type GetOrderHandler struct {
	OrderBook orderbook.OrderBook
}

func (handler *GetOrderHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	o := handler.OrderBook.GetOrderByHash(params["order"])

	if o == nil {
		http.NotFound(rw, r)
		return
	}

	json.NewEncoder(rw).Encode(o)
}
