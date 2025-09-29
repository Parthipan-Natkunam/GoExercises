package hyperfactorial

import (
	"math"
	"math/big"
)

func Hyperfactorial(n int) *big.Int {
	if n < 0 {
		panic("n must be a non-negative integer")
	}
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil))
	}
	return result
}

func HyperfactorialAsymptotic(n int) float64 {
	if n < 0 {
		panic("n must be a non-negative integer")
	}
	if n == 0 {
		return 1
	}
	const glaisherKinkelinConstant = 1.28242712910062261468
	firstExponent := float64(n*n + n) / 2 + 1.0/12.0
	secondExponent := float64(n*n)/4
	return math.Round(glaisherKinkelinConstant * math.Pow(float64(n), firstExponent) * math.Exp(-secondExponent))
}
