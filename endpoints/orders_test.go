package endpoints

import (
	"testing"

	"github.com/DexyProject/dexy-go/types"
)

var pricetests = []struct {
	expected string
	err      bool
	give     string
	get      string
	getEth   bool
}{
	{"", true, "fail", "fail", true},
	{"", true, "fail", "1", true},
	{"1", true, "3000000000000000000", "3000000000000000000", false},
	{"0.3333333333333333", true, "1000000000000000000", "3000000000000000000", false},
	{"0.0018499999999894799", true, "67489986216600", "124856474500", true},
}

func Test_CalculatePrice(t *testing.T) {

	for _, tt := range pricetests {
		order := types.Order{}

		order.Get.Amount = tt.get
		order.Give.Amount = tt.give

		if tt.getEth {
			order.Get.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Give.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		} else {
			order.Give.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Get.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		}

		price, err := calculatePrice(order)

		if err != nil && !tt.err {
			t.Errorf("error was not expected")
		}

		if price != tt.expected {
			t.Errorf("price %s did not match expected %s", price, tt.expected)
		}
	}
}
