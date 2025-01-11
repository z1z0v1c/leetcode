package countingbits

import (
	"slices"
	"testing"
)

func TestCountBits(t *testing.T) {
	// Example one
	n := 2
	expected := []int{0, 1, 1}
	actual := countBits(n)

	if slices.Equal(expected, actual) {
		t.Errorf("Expected and actual results are not equal.")
	}

	// Example two
	n = 5
	expected = []int{0, 1, 2, 1, 2}
	actual = countBits(n)

	if slices.Equal(expected, actual) {
		t.Errorf("Expected and actual results are not equal.")
	}
}