package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

type Trade struct {
	Token  common.Address `json:"token"`
	Amount string         `json:"amount"`
}

type Order struct {
	Hash      string         `json:"hash"`
	Price     float64        `json:"price"`
	Give      Trade          `json:"give"`
	Get       Trade          `json:"get"`
	Expires   string         `json:"expires"`
	Nonce     string         `json:"nonce"`
	User      common.Address `json:"user"`
	Exchange  common.Address `json:"exchange"`
	Signature EC             `json:"signature"`
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

	sha.Write(order.Get.Token[:])
	sha.Write(amountGet[:])
	sha.Write(order.Give.Token[:])
	sha.Write(amountGive[:])
	sha.Write(expires[:])
	sha.Write(nonce[:])
	sha.Write(order.User[:])
	sha.Write(order.Exchange[:])

	return sha.Sum(nil), nil
}
