package twosumiisortedarray

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// Example one
	numbers, target := []int{2,7,11,15}, 9

	expected := []int{1, 2}
	actual := twoSum(numbers, target)

	if !slices.Equal(expected, actual) {
		t.Errorf("twoSum(%v, %d) returned %v instead of %v.", numbers, target, actual, expected)
	}

	// Example two
	numbers, target = []int{2, 3, 4}, 6

	expected = []int{1, 3}
	actual = twoSum(numbers, target)

	if !slices.Equal(expected, actual) {
		t.Errorf("twoSum(%v, %d) returned %v instead of %v.", numbers, target, actual, expected)
	}

	// Example three
	numbers, target = []int{-1, 0}, -1

	expected = []int{1, 2}
	actual = twoSum(numbers, target)

	if !slices.Equal(expected, actual) {
		t.Errorf("twoSum(%v, %d) returned %v instead of %v.", numbers, target, actual, expected)
	}
}