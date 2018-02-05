package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EC struct {
	V        int   `json:"v" bson:"v"`
	R        Bytes `json:"r" bson:"r"`
	S        Bytes `json:"s" bson:"s"`
	Prefixed bool  `json:"prefixed" bson:"prefixed"`
}

func (ec *EC) Verify(address Address, hash Hash) bool {

	sigBytes := make([]byte, 65)
	copy(sigBytes[32-len(ec.R):32], ec.R[:])
	copy(sigBytes[64-len(ec.S):64], ec.S[:])
	sigBytes[64] = byte(ec.V - 27)

	var hashedBytes = hash[:]
	if ec.Prefixed {
		hashedBytes = crypto.Keccak256(append([]byte("\x19Ethereum Signed Message:\n32"), hash[:]...))
	}

	pub, err := crypto.Ecrecover(hashedBytes, sigBytes)
	if err != nil {
		return false
	}

	recoverAddress := common.BytesToAddress(crypto.Keccak256(pub[1:])[12:])
	return reflect.DeepEqual(address.Address[:], recoverAddress[:])
}
