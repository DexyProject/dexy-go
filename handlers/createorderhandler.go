package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/validators"
)

type CreateOrderHandler struct {
	OrderBook orderbook.OrderBook
	BalanceValidator *validators.BalanceValidatorSession
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

	err = handler.OrderBook.InsertOrder(o)
	if err != nil {
		return
	}

	_, err = handler.BalanceValidator.CheckBalance(o.Exchange, o.User)
	if err != nil {
		return
		// @todo
	}

	// @todo response
}
