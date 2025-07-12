package main

import (
	"reflect"
	"testing"
)


func TestCountCharFrequency(t *testing.T) {
	tests := []struct {
		input    string
		expected map[rune]int
	}{
		{
			input: "ddcc",
			expected: map[rune]int{
				'd': 2,
				'c': 2,
			},
		},
		{
			input: "Go! Go?",
			expected: map[rune]int{
				'G': 2,
				'o': 2,
			},
		},
		{
			input: "a\nb\rc",
			expected: map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
		},
	}

	for _, test := range tests {
		result := Counter(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %q, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
