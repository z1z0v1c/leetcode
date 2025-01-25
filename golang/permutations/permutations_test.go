package permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	// Example one
	nums := []int{1, 2, 3}

	expected := [][]int{{1,2,3},{1,3,2},{2,1,3},{2,3,1},{3,1,2},{3,2,1}}
	actual := permute(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("permute(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}

	// Example two
	nums = []int{0, 1}

	expected = [][]int{{0, 1},{1, 0}}
	actual = permute(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("permute(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}

	// Example three
	nums = []int{1}

	expected = [][]int{{1}}
	actual = permute(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("permute(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}
}