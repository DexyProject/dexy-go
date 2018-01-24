package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type History struct {
	History history.History
}

func (history *History) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")

	if token == "0x0000000000000000000000000000000000000000" || !common.IsHexAddress(token) {
		// @todo error body
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("user"))

	addr := types.HexToAddress(token)

	h := history.History.GetHistory(addr, user, limit)
	json.NewEncoder(rw).Encode(h)
}
