package main

import (
	"go_excercises/fireworks"
	"testing"
)

func TestGetFinaleStart(t *testing.T) {
	tests := []struct {
		name          string
		input         []fireworks.Firework
		expectedIndex int
	}{
		{
			name: "Test case 1",
			input: []fireworks.Firework{
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
			name: "Test case 2",
			input: []fireworks.Firework{
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
			name: "Test case 3",
			input: []fireworks.Firework{},
			expectedIndex: -1,
		},
		{
			name: "Test case 4",
			input: []fireworks.Firework{
				{Height: 10, Size: 6, Velocity: 4},
			},
			expectedIndex: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			startIndex := fireworks.GetFinaleStart(tc.input)
			if startIndex != tc.expectedIndex {
				t.Errorf("Expected finale start index %d, got %d", tc.expectedIndex, startIndex)
			}
		})
	}
}
