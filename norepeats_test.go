package main

import (
	"goexcercises/norepeats"
	"testing"
)

func TestNonRepeat(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLastChar rune
	}{
		{
			name: "Test case 1",
			input: "candy canes do taste yummy",
			expectedLastChar: 'u',
		},
		{
			name: "Test case 2",
			input: "hello world",
			expectedLastChar: 'd',
		},
		{
			name: "Test case 3",
			input: " ",
			expectedLastChar: ' ',
		},
		{
			name: "Test case 4",
			input: "aabbccddeeff gghhii",
			expectedLastChar: ' ',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lastUniqueChar := norepeats.NonRepeat(tc.input)
			if lastUniqueChar != tc.expectedLastChar {
				t.Errorf("Expected finale start index %d, got %d", tc.expectedLastChar, lastUniqueChar)
			}
		})
	}
}