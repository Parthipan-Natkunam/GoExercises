package hyperfactorial

import "math/big"

const NEGATIVE_INPUT_ERROR = "n must be a non-negative integer"

func Hyperfactorial(n int) *big.Int {
	if n < 0 {
		panic(NEGATIVE_INPUT_ERROR)
	}
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil))
	}
	return result
}
