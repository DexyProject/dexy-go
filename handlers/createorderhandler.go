package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/common"

)

type CreateOrderHandler struct {
	OrderBook orderbook.OrderBook
	BalanceValidator validators.BalanceValidator
}

func (handler *CreateOrderHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var o types.Order
	err := decoder.Decode(o)
	if err != nil {
		// @todo
		return
	}

	ok, err := handler.BalanceValidator.CheckBalance(o)
	if !ok {
		// @todo
		return
	}

	hash, err := o.OrderHash()
	if err != nil {
		// @todo
		return
	}

	o.Hash = common.ToHex(hash)
	price, err := calculatePrice(o)
	if err != nil {
		// @todo
		return
	}

	o.Price = price
	err = handler.OrderBook.InsertOrder(o)
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
