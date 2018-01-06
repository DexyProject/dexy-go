package types

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func StringToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(hexString, "0x"))
}

func IntStringToBytes(intString string) ([]byte, error) {
	bigInt := new(big.Int)
	_, success := bigInt.SetString(intString, 10)
	if !success {
		return nil, errors.New("value not a valid integer")
	}

	return abi.U256(bigInt), nil
}
