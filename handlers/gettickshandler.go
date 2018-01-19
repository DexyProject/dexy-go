package handlers

import (
	"net/http"
	"github.com/DexyProject/dexy-go/db"
	"encoding/json"
)

type GetTicksHandler struct {
	TickQuery db.TickQuery
}

func (handler *GetTicksHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	block := query.Get("block")

	h, err := handler.TickQuery.AggregateTick(block)
	if err != nil {
		// @todo error handling
		rw.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(rw).Encode(h)
}