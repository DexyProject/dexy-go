package types

import (
	"bytes"
	"testing"
)

func TestOrder_OrderHash(t *testing.T) {

	o := Order{
		Price:    "0.1",
		Give:     Trade{Token: HexToAddress("0x0000000000000000000000000000000000000000"), Amount: NewInt(30)},
		Get:      Trade{Token: HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: NewInt(300)},
		Expires:  1514892553,
		Nonce:    123,
		User:     HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Exchange: HexToAddress("0x58e91b0734e2b33efc86067ce4db128366f30dc9"),
	}

	expected := NewHash("0x56209e5e80fab187438ab24ba6d3df31b8c369b0e86d7b210358efd92eef1cfa")
	hash := o.OrderHash()

	if !bytes.Equal(expected.Hash[:], hash.Hash[:]) {
		t.Error("order hashes were not equal")
	}
}
