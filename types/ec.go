package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EC struct {
	V int    `json:"v" bson:"v"`
	R string `json:"r" bson:"r"`
	S string `json:"s" bson:"s"`
}

func (ec *EC) Verify(address common.Address, hash []byte) bool {

	r, err := StringToBytes(ec.R)
	if err != nil {
		return false
	}

	s, err := StringToBytes(ec.S)
	if err != nil {
		return false
	}

	sigBytes := make([]byte, 65)
	copy(sigBytes[32-len(r):32], r[:])
	copy(sigBytes[64-len(s):64], s[:])
	sigBytes[64] = byte(ec.V - 27)

	hashedBytes := append([]byte("\x19Ethereum Signed Message:\n32"), hash[:]...)
	signedBytes := crypto.Keccak256(hashedBytes)
	pub, err := crypto.Ecrecover(signedBytes, sigBytes)
	if err != nil {
		return false
	}

	recoverAddress := common.BytesToAddress(crypto.Keccak256(pub[1:])[12:])
	return reflect.DeepEqual(address[:], recoverAddress[:])
}
