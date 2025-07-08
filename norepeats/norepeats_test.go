package norepeats

import (
	"testing"
)

func TestNonRepeat(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLastChar rune
	}{
		{
			name: "Base Case",
			input: "candy canes do taste yummy",
			expectedLastChar: 'u',
		},
		{
			name: "Last Character",
			input: "hello world",
			expectedLastChar: 'd',
		},
		{
			name: "Empty String",
			input: " ",
			expectedLastChar: ' ',
		},
		{
			name: "No Unique Characters",
			input: "aabbccddeeff gghhii",
			expectedLastChar: ' ',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lastUniqueChar := NonRepeat(tc.input)
			if lastUniqueChar != tc.expectedLastChar {
				t.Errorf("Expected last non-repat character '%c', got '%c'", tc.expectedLastChar, lastUniqueChar)
			}
		})
	}
}