package endpoints

import (
	"testing"

	"github.com/DexyProject/dexy-go/types"
)

var pricetests = []struct {
	expected float64
	err      bool
	give     int64
	get      int64
	getEth   bool
}{
	{1.00000000, false, 3000000000000000000, 3000000000000000000, false},
	{0.3333333333333333, false, 1000000000000000000, 3000000000000000000, false},
	{0.0018499999999894799, false, 67489986216600, 124856474500, true},
	{0, true, 0, 124856474500, true},
}

func Test_CalculatePrice(t *testing.T) {

	for _, tt := range pricetests {
		order := types.Order{}

		order.Take.Amount = types.NewInt(tt.get)
		order.Make.Amount = types.NewInt(tt.give)

		if tt.getEth {
			order.Take.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Make.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		} else {
			order.Make.Token = types.HexToAddress("0x0000000000000000000000000000000000000000")
			order.Take.Token = types.HexToAddress("0xaaa21488d380648c240a6444996b8ee81fb5b762")
		}

		price, err := calculatePrice(order)

		if !tt.err && err != nil {
			t.Errorf("error was not expected")
		}

		if price != tt.expected {
			t.Errorf("price %s did not match expected %s", price, tt.expected)
		}
	}
}
