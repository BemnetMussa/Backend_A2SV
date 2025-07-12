package main

import (
	"testing"
)

type testType struct {
	input string
	expected bool
}

func TestPalindromeCheck(t *testing.T) {
	tests := []testType {
		{"racecar", true},
		{"RaceCar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"No lemon, no melon", true},
		{"Was it a car or a cat I saw?", true},
		{"hello world", false},
		{"", true}, // empty strings are palindrome
	}

	for _, test := range tests {
		result := PalindromeCheck(test.input)
		if result != test.expected {
			t.Errorf("palindromeCheck(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}