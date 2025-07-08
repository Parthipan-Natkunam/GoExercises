package fireworks

import (
	"testing"
)

func TestGetFinaleStart(t *testing.T) {
	tests := []struct {
		name          string
		input         []Firework
		expectedIndex int
	}{
		{
			name: "Base Case",
			input: []Firework{
				{Height: 10, Size: 6, Velocity: 4},
				{Height: 13, Size: 3, Velocity: 2},
				{Height: 17, Size: 6, Velocity: 3},
				{Height: 21, Size: 8, Velocity: 4},
				{Height: 19, Size: 5, Velocity: 3},
				{Height: 18, Size: 4, Velocity: 4},
			},
			expectedIndex: 2,
		},
		{
			name: "Multiple Subsequences With Varying Lengths",
			input: []Firework{
				{Height: 10, Size: 6, Velocity: 4},
				{Height: 13, Size: 3, Velocity: 2},
				{Height: 17, Size: 6, Velocity: 3},
				{Height: 21, Size: 8, Velocity: 4},
				{Height: 19, Size: 5, Velocity: 3},
				{Height: 18, Size: 4, Velocity: 4},
				{Height: 31, Size: 3, Velocity: 2},
				{Height: 17, Size: 6, Velocity: 3},
				{Height: 21, Size: 8, Velocity: 4},
				{Height: 19, Size: 5, Velocity: 3},
				{Height: 18, Size: 4, Velocity: 4},
				{Height: 17, Size: 6, Velocity: 3},
				{Height: 21, Size: 8, Velocity: 4},
				{Height: 19, Size: 5, Velocity: 3},
				{Height: 18, Size: 4, Velocity: 4},
			},
			expectedIndex: 7,
		},
		{
			name:          "Empty Input",
			input:         []Firework{},
			expectedIndex: -1,
		},
		{
			name: "Single Firework",
			input: []Firework{
				{Height: 10, Size: 6, Velocity: 4},
			},
			expectedIndex: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			startIndex := GetFinaleStart(tc.input)
			if startIndex != tc.expectedIndex {
				t.Errorf("Expected finale start index %d, got %d", tc.expectedIndex, startIndex)
			}
		})
	}
}
