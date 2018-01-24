package endpoints

import (
	"testing"

	"github.com/DexyProject/dexy-go/types"
	"math/big"
)

var pricetests = []struct {
	expected string
	err      bool
	give     int64
	get      int64
	getEth   bool
}{
	{"1", true, 3000000000000000000, 3000000000000000000, false},
	{"0.3333333333", true, 1000000000000000000, 3000000000000000000, false},
	{"0.00185", true, 67489986216600, 124856474500, true},
}

func Test_CalculatePrice(t *testing.T) {

	for _, tt := range pricetests {
		order := types.Order{}

		order.Get.Amount = *new(big.Int).SetInt64(tt.get)
		order.Give.Amount = *new(big.Int).SetInt64(tt.give)

		if tt.getEth {
			order.Get.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Give.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		} else {
			order.Give.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Get.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		}

		price := calculatePrice(order)

		if price != tt.expected {
			t.Errorf("price %s did not match expected %s", price, tt.expected)
		}
	}
}
