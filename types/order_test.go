package types

import (
	"bytes"
	"testing"
	"time"
)

func TestOrder_OrderHash(t *testing.T) {

	o := Order{
		Price:    0.1,
		Give:     Trade{Token: HexToAddress("0x0000000000000000000000000000000000000000"), Amount: NewInt(1000000000000000)},
		Get:      Trade{Token: HexToAddress("0xbebb2325ef529e4622761498f1f796d262100768"), Amount: NewInt(10000000)},
		Expires:  Timestamp{time.Unix(1519216353, 0)},
		Nonce:    1519216151661,
		User:     HexToAddress("0x3b6760e4bae3d347adaf6a36523f901bbd7ed7f1"),
		Exchange: HexToAddress("0x3db7a4c4c30eaec1ac7301c3f95920afbe6719e3"),
	}

	expected := NewHash("0x69d933eba45973031c42b55b432614c9ae024e69f65c04b3badc7baf22275a9e")
	hash := o.OrderHash()

	if !bytes.Equal(expected.Hash[:], hash.Hash[:]) {
		t.Error("order hashes were not equal")
	}
}
