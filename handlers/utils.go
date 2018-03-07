package handlers

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func getLimit(limit string) int {
	if len(limit) != 0 && limit != "0" {

		u, err := strconv.Atoi(limit)
		if err == nil {
			return u
		}
	}

	return 100
}

func getUser(user string) *common.Address {
	if user == "" || !common.IsHexAddress(user) {
		return nil
	}

	addr := common.HexToAddress(user)
	return &addr
}
