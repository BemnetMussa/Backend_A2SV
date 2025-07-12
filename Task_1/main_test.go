package main

import (
	"testing"
)
func TestAverageGradeCalculator(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected float32
	}{
		{"All 100s", map[string]int{"Math": 100, "CS": 100}, 100.0},
		{"Mixed Grades", map[string]int{"Math": 80, "Eng": 90}, 85.0},
		{"One Grade", map[string]int{"Art": 70}, 70.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			profile = test.input // set global variable
			got := averageGradeCalculator()
			if got != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, got)
			}
		})
	}
}
