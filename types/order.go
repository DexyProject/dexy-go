package types

import (
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Trade struct {
	Token  Address `json:"token" bson:"token"`
	Amount Int     `json:"amount" bson:"amount"`
}

type Orders struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

type Order struct {
	Hash      Hash  `json:"hash,omitempty" bson:"hash"`
	Price     string  `json:"price,omitempty" bson:"price"`
	Give      Trade   `json:"give" bson:"give"`
	Get       Trade   `json:"get" bson:"get"`
	Expires   int64   `json:"expires" bson:"expires"`
	Nonce     int64   `json:"nonce" bson:"nonce"`
	User      Address `json:"user" bson:"user"`
	Exchange  Address `json:"exchange" bson:"exchange"`
	Signature EC      `json:"signature" bson:"signature"`
}

func (o *Order) OrderHash() Hash {
	if o.Hash.String() == (Hash{}).String() {
		o.generateHash()
	}

	return o.Hash
}

func (o *Order) generateHash() {
	sha := sha3.NewKeccak256()

	sha.Write(o.Get.Token.Address[:])
	sha.Write(o.Get.Amount.U256()[:])
	sha.Write(o.Give.Token.Address[:])
	sha.Write(o.Give.Amount.U256()[:])
	sha.Write(NewInt(o.Expires).U256()[:])
	sha.Write(NewInt(o.Nonce).U256()[:])
	sha.Write(o.User.Address[:])
	sha.Write(o.Exchange.Address[:])

	o.Hash.SetBytes(sha.Sum(nil))
}
