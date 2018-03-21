package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DexyProject/dexy-go/history"
	dexyhttp "github.com/DexyProject/dexy-go/http"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type History struct {
	History history.History
}

func (history *History) Handle(rw http.ResponseWriter, r *http.Request) error {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("user"))

	addr := types.HexToAddress(token)

	h := history.History.GetHistory(addr, user, limit)
	json.NewEncoder(rw).Encode(h)
	return nil
}
