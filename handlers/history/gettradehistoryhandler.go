package history

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/handlers"
	"github.com/DexyProject/dexy-go/history"
	"github.com/ethereum/go-ethereum/common"
)

type GetTradeHistoryHandler struct {
	History history.History
}

func (handler *GetTradeHistoryHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")

	if token == "0x0000000000000000000000000000000000000000" || !common.IsHexAddress(token) {
		// @todo error body
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := handlers.GetLimit(query.Get("limit"))
	user := handlers.GetUser(query.Get("user"))

	addr := common.HexToAddress(token)

	h := handler.History.GetHistory(addr, user, limit)
	json.NewEncoder(rw).Encode(h)
}
