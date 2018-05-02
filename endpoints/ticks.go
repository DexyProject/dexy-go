package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	dexyhttp "github.com/DexyProject/dexy-go/http"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
	"go.uber.org/zap"
)

type Ticks struct {
	Ticks ticks.Ticks
}

func (ep *Ticks) GetTicks(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	token := query.Get("token")

	if token == types.ETH_ADDRESS {
		return dexyhttp.NewError(fmt.Sprintf("invalid token: %s", types.ETH_ADDRESS), http.StatusBadRequest)
	}

	addr := types.HexToAddress(token)
	h, err := ep.Ticks.FetchTicks(addr)
	if err != nil {
		log.Error("could not fetch ep", zap.Error(err))
		return dexyhttp.NewError("fetching ep failed", http.StatusBadRequest)
	}

	json.NewEncoder(rw).Encode(h)
	return nil
}
