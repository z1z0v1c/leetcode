package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	// Example one
	target := 7
	candidates := []int{2, 3, 6, 7}

	expected := [][]int{{2, 2, 3}, {7}}
	actual := combinationSum(candidates, target)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combinationSum(%#v, %d) returned %#v instead of %#v.", candidates, target, actual, expected)
	}

	// Example two
	target = 8
	candidates = []int{2, 3, 5}

	expected = [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	actual = combinationSum(candidates, target)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combinationSum(%#v, %d) returned %#v instead of %#v.", candidates, target, actual, expected)
	}

	// Example three
	target = 1
	candidates = []int{2}

	expected = nil
	actual = combinationSum(candidates, target)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combinationSum(%#v, %d) returned %#v instead of %#v.", candidates, target, actual, expected)
	}
}
