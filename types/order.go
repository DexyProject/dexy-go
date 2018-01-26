package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	Hash      Hash    `json:"hash,omitempty" bson:"hash"`
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

func (order *Order) Validate() error {
	if !common.IsHexAddress(order.Get.Token.String()) || !common.IsHexAddress(order.Give.Token.String()) {
		return fmt.Errorf("address is not valid hex")
	}

	if strings.ToLower(order.Give.Token.String()) == strings.ToLower(order.Get.Token.String()) {
		return fmt.Errorf("token addresses are identical")
	}

	if order.Expires <= time.Now().Unix() {
		return fmt.Errorf("invalid expires time: %v", order.Expires)
	}

	return nil
}
