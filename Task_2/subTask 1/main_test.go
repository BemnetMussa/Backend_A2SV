package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	input := "Hello, hello world! Hello?"
	expected := map[string]int{
		"hello": 3,
		"world": 1,
		"interest": 1,
	}

	result := WordFrequency(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
	