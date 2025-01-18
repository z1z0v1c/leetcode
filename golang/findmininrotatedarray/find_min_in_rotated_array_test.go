package findmininrotatedarray

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	// Example one
	nums := []int{3,4,5,1,2}

	expected := 1
	actual := findMin(nums)

	if expected != actual {
		t.Errorf("findMin(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example two
	nums = []int{4,5,6,7,0,1,2}

	expected = 0
	actual = findMin(nums)

	if expected != actual {
		t.Errorf("findMin(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example three
	nums = []int{11,13,15,17}

	expected = 11
	actual = findMin(nums)

	if expected != actual {
		t.Errorf("findMin(%#v) returned %d instead of %d.", nums, actual, expected)
	}
}
 