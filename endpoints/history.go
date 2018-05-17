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
	quote := query.Get("quote")
	base := query.Get("base")

	// @todo check if base is valid

	if quote == types.ETH_ADDRESS || !common.IsHexAddress(quote) {
		return dexyhttp.NewError(fmt.Sprintf("invalid quote: %s", quote), http.StatusBadRequest)
	}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("maker"))

	addr := types.HexToAddress(quote)

	h := ep.History.GetHistory(addr, user, limit)
	json.NewEncoder(rw).Encode(h)
	return nil
}
