package endpoints

import (
	"encoding/json"
	"net/http"

	"log"

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
	if token == types.ETH_ADDRESS {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	addr := types.HexToAddress(token)
	h, err := ticks.Ticks.FetchTicks(addr)
	if err != nil {
		log.Printf("could not fetch ticks: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(rw).Encode(h)
}
