package combinationsumii

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	// Example one
	target := 8
	candidates := []int{10, 1, 2, 7, 6, 1, 5}

	expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	actual := combinationSum2(candidates, target)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combinationSum2(%#v, %d) returned %#v instead of %#v.", candidates, target, actual, expected)
	}

	// Example two
	target = 5
	candidates = []int{2, 5, 2, 1, 2}

	expected = [][]int{{1, 2, 2}, {5}}
	actual = combinationSum2(candidates, target)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combinationSum2(%#v, %d) returned %#v instead of %#v.", candidates, target, actual, expected)
	}
}
