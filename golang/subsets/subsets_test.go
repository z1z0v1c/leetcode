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

	// Example three
	nums = []int{9,0,3,5,7}
	
	expected = [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}
	actual = subsets(nums)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("subsets(%#v) returned %#v instead of %#v.", nums, actual, expected)
	}
}