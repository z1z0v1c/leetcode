package subsets

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	// Example one
	nums := []int{1, 2, 3}
	
	expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	actual := subsets(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsets(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}

	// Example two
	nums = []int{0}
	
	expected = [][]int{{}, {0}}
	actual = subsets(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsets(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}
}