package types

import (
	"testing"
)

func TestEC_Verify(t *testing.T) {
	address := HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
	hash, err := StringToBytes("0x1f406680ed210e1589ad11769afa120d7f171ce1daf91ea3346b189ce9935203")
	if err != nil {
		t.Fail()
	}

	sig := EC{
		V: 28,
		R: "0x490149269faa99a814bdb4c34be3205ebbdfcbc9079c6c3779d3ea80f1d06f8d",
		S: "0x28eec0e2deff6903087d597bf6f728cf30d333b6efb5405d7e0c555e6964608b",
	}

	if !sig.Verify(address, hash) {
		t.Error("failed to verify signature")
	}
}
