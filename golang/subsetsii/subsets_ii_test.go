package subsetsii

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	// Example one
	nums := []int{1, 2, 2}

	expected := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	actual := subsetsWithDup(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsetsWithDup(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}

	// Example two
	nums = []int{0}

	expected = [][]int{{}, {0}}
	actual = subsetsWithDup(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsetsWithDup(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}

	// Example two
	nums = []int{4, 4, 4, 1, 4}

	expected = [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}}
	actual = subsetsWithDup(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsetsWithDup(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}
}
