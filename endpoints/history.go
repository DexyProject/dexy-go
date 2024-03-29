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

func (ep *History) Handle(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS || !common.IsHexAddress(token) {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", token), http.StatusBadRequest)
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("maker"))

	addr := types.HexToAddress(token)

	h := ep.History.GetHistory(addr, user, limit)
	json.NewEncoder(rw).Encode(h)
	return nil
}
