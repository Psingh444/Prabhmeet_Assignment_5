package main

import (
	"testing"
)

// TestSubtract tests the Subtract function with various inputs and expected outputs.
func TestSubtract(t *testing.T) {
	tests := []struct {
		minuend, subtrahend int
		expectedDifference  int
	}{
		{10, 5, 5},
		{0, 0, 0},
		{-5, -5, 0},
		{5, 10, -5},
		{10, -5, 15},
		{-10, 5, -15},
	}

	for _, test := range tests {
		result := Subtract(test.minuend, test.subtrahend)
		if result != test.expectedDifference {
			t.Errorf("Subtract(%d, %d) = %d; expected %d", test.minuend, test.subtrahend, result, test.expectedDifference)
		}
	}
}
