package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			"case 1",
			112,
			false,
		},
		{
			"case 2",
			11011,
			true,
		},
		{
			"case 3",
			2332,
			true,
		},
		{
			"case 4",
			1234,
			false,
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.input)
		if got != test.expected {
			t.Errorf("expected %t but got %t", test.expected, got)
		}
	}
}
