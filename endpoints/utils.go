package endpoints

import (
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
