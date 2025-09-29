package hyperfactorial

import (
	"math/big"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name          string
		input         int
		expected      *big.Int
		shouldPanic   bool
	}{
		{name: "Zero", input: 0, expected: big.NewInt(1), shouldPanic: false},
		{name: "One", input: 1, expected: big.NewInt(1), shouldPanic: false},
		{name: "Two", input: 2, expected: big.NewInt(4), shouldPanic: false},
		{name: "Three", input: 3, expected: big.NewInt(108), shouldPanic: false},
		{name: "Seven", input: 7, expected: big.NewInt(3319766398771200000), shouldPanic: false},
		{name:"Negative", input: -1, expected: nil, shouldPanic: true}, 
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for negative input")
					}
				}()
			}
			result := Hyperfactorial(tc.input)
			if !tc.shouldPanic && result.Cmp(tc.expected) != 0 {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}


