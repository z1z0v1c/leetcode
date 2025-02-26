package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	// Exmple one
	nums := []int{2, 3, -2, 4}

	expected := 6
	actual := maxProduct(nums)
	
	if expected != actual {
		t.Errorf("maxProduct(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Exmple two
	nums = []int{-2, 0, -1}

	expected = 0
	actual = maxProduct(nums)
	
	if expected != actual {
		t.Errorf("maxProduct(%#v) returned %d instead of %d.", nums, actual, expected)
	}
}