package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Trade struct {
	Token  Address `json:"token" bson:"token"`
	Amount big.Int `json:"amount" bson:"amount"`
}

type Orders struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

type Order struct {
	Hash      string  `json:"hash,omitempty" bson:"hash"`
	Price     string  `json:"price,omitempty" bson:"price"`
	Give      Trade   `json:"give" bson:"give"`
	Get       Trade   `json:"get" bson:"get"`
	Expires   string  `json:"expires" bson:"expires"`
	Nonce     string  `json:"nonce" bson:"nonce"`
	User      Address `json:"user" bson:"user"`
	Exchange  Address `json:"exchange" bson:"exchange"`
	Signature EC      `json:"signature" bson:"signature"`
}

func (order *Order) OrderHash() ([]byte, error) {
	sha := sha3.NewKeccak256()

	expires, err := IntStringToBytes(order.Expires)
	if err != nil {
		return nil, err
	}

	amountGive := abi.U256(&order.Give.Amount)
	amountGet := abi.U256(&order.Get.Amount)

	nonce, err := IntStringToBytes(order.Nonce)
	if err != nil {
		return nil, err
	}

	sha.Write(order.Get.Token.Address[:])
	sha.Write(amountGet[:])
	sha.Write(order.Give.Token.Address[:])
	sha.Write(amountGive[:])
	sha.Write(expires[:])
	sha.Write(nonce[:])
	sha.Write(order.User.Address[:])
	sha.Write(order.Exchange.Address[:])

	return sha.Sum(nil), nil
}
