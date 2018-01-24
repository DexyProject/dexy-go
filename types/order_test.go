package types

import (
	"bytes"
	"testing"
)

func TestOrder_OrderHash(t *testing.T) {

	o := Order{
		Hash:     "0xcd09064f280940d69e0d3d9741c7986b80935dbe066a218cc83823e7eb518681",
		Price:    "0.1",
		Give:     Trade{Token: HexToAddress("0x0000000000000000000000000000000000000000"), Amount: "30"},
		Get:      Trade{Token: HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: "300"},
		Expires:  "1514892553",
		Nonce:    "123",
		User:     HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Exchange: HexToAddress("0x58e91b0734e2b33efc86067ce4db128366f30dc9"),
	}

	hashed, err := o.OrderHash()
	if err != nil {
		t.Error(err)
	}

	byteHash, err := StringToBytes(o.Hash)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(hashed, byteHash) {
		t.Error("order hashes were not equal")
	}
}
