package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
	"log"
)

type Ticks struct {
	Ticks ticks.Ticks
}

func (ticks *Ticks) GetTicks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	token := query.Get("token")
	if token == types.ETH_ADDRESS {
		log.Printf("token request is 0x0")
		rw.WriteHeader(http.StatusBadRequest)
	}

	addr := types.HexToAddress(token)
	h, err := ticks.Ticks.FetchTicks(addr)
	if err != nil {
		log.Printf("could not fetch ticks: %v", err)
		rw.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(rw).Encode(h)
}
