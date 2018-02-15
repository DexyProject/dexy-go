package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
)

type Ticks struct {
	Ticks ticks.Ticks
}

func (ticks *Ticks) GetTicks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")
	addr := types.HexToAddress(token)

	h, err := ticks.Ticks.FetchTicks(addr)
	if err != nil {
		// @todo error handling
		rw.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(rw).Encode(h)
}
