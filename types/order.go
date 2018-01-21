package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Trade struct {
	Token  string `json:"token" bson:"token"`
	Amount string `json:"amount" bson:"amount"`
}

type Orders struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

type Order struct {
	Hash      string         `json:"hash,omitempty" bson:"hash"`
	Price     string         `json:"price,omitempty" bson:"price"`
	Give      Trade          `json:"give" bson:"give"`
	Get       Trade          `json:"get" bson:"get"`
	Expires   string         `json:"expires" bson:"expires"`
	Nonce     string         `json:"nonce" bson:"nonce"`
	User      string `json:"user" bson:"user"`
	Exchange  string `json:"exchange" bson:"exchange"`
	Signature EC             `json:"signature" bson:"signature"`
}

func (order *Order) OrderHash() ([]byte, error) {
	sha := sha3.NewKeccak256()

	expires, err := IntStringToBytes(order.Expires)
	if err != nil {
		return nil, err
	}

	amountGive, err := IntStringToBytes(order.Give.Amount)
	if err != nil {
		return nil, err
	}

	amountGet, err := IntStringToBytes(order.Get.Amount)
	if err != nil {
		return nil, err
	}

	nonce, err := IntStringToBytes(order.Nonce)
	if err != nil {
		return nil, err
	}

	get := common.HexToAddress(order.Get.Token)
	give := common.HexToAddress(order.Give.Token)
	userAddr := common.HexToAddress(order.User)
	exchangeAddr := common.HexToAddress(order.Exchange)

	sha.Write(get[:])
	sha.Write(amountGet[:])
	sha.Write(give[:])
	sha.Write(amountGive[:])
	sha.Write(expires[:])
	sha.Write(nonce[:])
	sha.Write(userAddr[:])
	sha.Write(exchangeAddr[:])

	return sha.Sum(nil), nil
}
