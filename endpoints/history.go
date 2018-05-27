package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/types"
)

type History struct {
	History history.History
}

func (ep *History) Handle(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	quote := query.Get("quote")
	base := query.Get("base")

	// @todo check if base is valid
	//if quote == types.ETH_ADDRESS || !common.IsHexAddress(quote) {
	//	return dexyhttp.NewError(fmt.Sprintf("invalid quote: %s", quote), http.StatusBadRequest)
	//}

	limit := GetLimit(query.Get("limit"))
	user := GetUser(query.Get("maker"))



	h := ep.History.GetHistory(
		types.Pair{Quote: types.HexToAddress(quote), Base: types.HexToAddress(base)},
		user,
		limit,
	)

	json.NewEncoder(rw).Encode(h)
	return nil
}
