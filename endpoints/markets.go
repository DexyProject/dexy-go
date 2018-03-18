package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
)

type Markets struct {
	OrderBook orderbook.OrderBook
}

func (m *Markets) GetMarkets(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tokens := make([]types.Address, 0)
	err := json.NewDecoder(r.Body).Decode(&tokens)
	defer r.Body.Close()

	if err != nil {
		log.Printf("unmarshalling json failed: %v", err.Error())
		returnError(rw, "badly formatted token list", http.StatusBadRequest)
		return
	}

	if len(tokens) == 0 {
		returnError(rw, "no tokens provided", http.StatusBadRequest)
		return
	}

	markets, err := m.OrderBook.GetMarkets(tokens)
	if err != nil {
		returnError(rw, "error fetching markets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(markets)
}
