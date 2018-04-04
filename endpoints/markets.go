package endpoints

import (
	"encoding/json"
	"net/http"

	dexyhttp "github.com/DexyProject/dexy-go/http"
	"github.com/DexyProject/dexy-go/markets"
	"github.com/DexyProject/dexy-go/types"
)

type Markets struct {
	Markets markets.Markets
}

func (m *Markets) GetMarkets(rw http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("tokens")

	tokens := make([]types.Address, 0)
	err := json.Unmarshal([]byte(query), &tokens)

	if err != nil {
		return dexyhttp.NewError("badly formatted token list", http.StatusBadRequest)
	}

	if len(tokens) == 0 {
		return dexyhttp.NewError("no tokens provided", http.StatusBadRequest)
	}

	ms, err := m.Markets.GetMarkets(tokens)
	if err != nil {
		return dexyhttp.NewError("error fetching Markets", http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(ms)
	return nil
}
