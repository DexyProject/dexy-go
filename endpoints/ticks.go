package endpoints

import (
	"net/http"
	"github.com/DexyProject/dexy-go/db"
	"encoding/json"
)

type Ticks struct {
	TickQuery db.TickQuery
}

func (ticks *Ticks) GetTicks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	block := query.Get("block")

	h, err := ticks.TickQuery.AggregateTick(block) //Output of AggregateTick is bson.M[]
	if err != nil {
		// @todo error handling
		rw.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(rw).Encode(h)
}
