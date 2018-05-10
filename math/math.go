package math

import (
	"math"
	"math/big"

	"github.com/DexyProject/dexy-go/types"
)

func ToUnitAmount(amount types.Int, decimals uint8) float64 {
	return ToUnitAmountForFloat(*new(big.Float).SetInt(&amount.Int), decimals)
}

func ToUnitAmountForFloat(amount big.Float, decimals uint8) float64 {
	a, _ := amount.Float64()
	return a / math.Pow(10.0, float64(decimals))
}
