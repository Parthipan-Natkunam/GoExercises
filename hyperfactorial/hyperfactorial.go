package hyperfactorial

import "math/big"

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
