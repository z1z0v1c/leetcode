package findtheduplicatenumber

import "testing"

func TestFindDuplicate(t *testing.T) {
	// Example one
	nums := []int{1, 3, 4, 2, 2}

	expected := 2
	actual := findDuplicate(nums)

	if expected != actual {
		t.Errorf("findDuplicate(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example two
	nums = []int{3, 1, 3, 4, 2}

	expected = 3
	actual = findDuplicate(nums)

	if expected != actual {
		t.Errorf("findDuplicate(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example three
	nums = []int{3, 3, 3, 3, 3}

	expected = 3
	actual = findDuplicate(nums)

	if expected != actual {
		t.Errorf("findDuplicate(%#v) returned %d instead of %d.", nums, actual, expected)
	}
}
