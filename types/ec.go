package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SigMode int

const (
	TYPED_SIG_EIP SigMode = 0
	GETH                  = 1
	TREZOR                = 2
)

type EC struct {
	V       int     `json:"v" bson:"v"`
	R       string  `json:"r" bson:"r"`
	S       string  `json:"s" bson:"s"`
	SigMode SigMode `json:"sig_mode" bson:"sig_mode"`
}

func (ec *EC) Verify(address Address, hash Hash) bool {

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

	hashBytes := getMessage(hash, ec.SigMode)
	if hashBytes == nil {
		return false
	}

	pub, err := crypto.Ecrecover(hashBytes, sigBytes)
	if err != nil {
		return false
	}

	recoverAddress := common.BytesToAddress(crypto.Keccak256(pub[1:])[12:])
	return reflect.DeepEqual(address.Address[:], recoverAddress[:])
}

// @todo having this here is pretty ugly
func getMessage(hash Hash, mode SigMode) []byte {
	switch mode {
	case TYPED_SIG_EIP:
		return hash[:]
	case GETH:
		return crypto.Keccak256(append([]byte("\x19Ethereum Signed Message:\n32"), hash[:]...))
	case TREZOR:
		return crypto.Keccak256(append([]byte("\x19Ethereum Signed Message:\n"), hash[:]...))
	}

	return nil
}
