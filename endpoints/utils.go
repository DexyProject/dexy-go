package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/common"
)

func GetLimit(limit string) int {
	if len(limit) != 0 && limit != "0" {

		u, err := strconv.Atoi(limit)
		if err == nil {
			return u
		}
	}

	return 100
}

func GetUser(user string) *types.Address {
	if user == "" || !common.IsHexAddress(user) {
		return nil
	}

	addr := types.HexToAddress(user)
	return &addr
}

func returnError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"error\": \"%v\"}", err)))
}
