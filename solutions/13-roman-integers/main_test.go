package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "case 1",
			input:    "III",
			expected: 3,
		},
		{
			name:     "case 2",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "case 3",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, test := range tests {
		got := romanToInt(test.input)
		if got != test.expected {
			t.Errorf("expected to get %d but got %d", test.expected, got)
		}
	}
}
