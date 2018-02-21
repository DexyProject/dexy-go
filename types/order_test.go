package types

import (
	"bytes"
	"testing"
)

func TestOrder_OrderHash(t *testing.T) {

	o := Order{
		Price:    "0.1",
		Give:     Trade{Token: HexToAddress("0x0000000000000000000000000000000000000000"), Amount: NewInt(1000000000000000)},
		Get:      Trade{Token: HexToAddress("0xbebb2325ef529e4622761498f1f796d262100768"), Amount: NewInt(100000000)},
		Expires:  UnixMilli(1519214183723),
		Nonce:    1519213982123,
		User:     HexToAddress("0x3b6760E4bAE3D347adAF6A36523F901bBD7Ed7f1"),
		Exchange: HexToAddress("0x3db7a4c4c30eaec1ac7301c3f95920afbe6719e3"),
	}

	expected := NewHash("0xb948dce070ec8d4a8d7975c7c090cd5a47befc91b3609b231bcd7508719538cf")
	hash := o.OrderHash()

	if !bytes.Equal(expected[:], hash[:]) {
		t.Error("order hashes were not equal")
	}
}
