package history

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/history"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
)

type GetTradeHistoryHandler struct {
	History history.History
}

func (handler *GetTradeHistoryHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	addr := common.HexToAddress(params["token"])

	h := handler.History.GetHistory(addr, nil, 10)
	json.NewEncoder(rw).Encode(h)

}
