package topkfrequentelements

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	// Example one
	nums, k := []int{1, 1, 1, 2, 2, 3}, 2

	expected := []int{1, 2}
	actual := topKFrequent(nums, k)

	if !slices.Equal(expected, actual) {
		t.Error("topKFrequent(nums, k) returned incorrect result in example 1.")
	}

	// Example two
	nums, k = []int{1}, 1

	expected = []int{1}
	actual = topKFrequent(nums, k)

	if !slices.Equal(expected, actual) {
		t.Error("topKFrequent(nums, k) returned incorrect result in example 1.")
	}
}