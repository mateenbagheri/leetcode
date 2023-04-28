package main

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "case 1",
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "case 2",
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "case 3",
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		got1 := twoSumsOptimal(test.input, test.target)
		got2 := twoSum(test.input, test.target)
		if !reflect.DeepEqual(test.expected, got1) {
			t.Errorf("expected %v but instead got %v", test.expected, got1)
		}
		if !reflect.DeepEqual(test.expected, got2) {
			t.Errorf("expected %v but instead got %v", test.expected, got2)
		}
	}
}
