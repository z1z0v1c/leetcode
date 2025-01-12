package missingnumber

import "testing"

func TestMissingNumber(t *testing.T) {
	// Example one
	nums := []int{3, 0, 1}
	expected := 2
	actual := missingNumber(nums)

	if expected != actual {
		t.Errorf("missingNumber(nums) returned %d instead of %d.", actual, expected)
	}
	
	// Example two
	nums = []int{0, 1}
	expected = 2
	actual = missingNumber(nums)

	if expected != actual {
		t.Errorf("missingNumber(nums) returned %d instead of %d.", actual, expected)
	}

	// Example three
	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected = 8
	actual = missingNumber(nums)

	if expected != actual {
		t.Errorf("missingNumber(nums) returned %d instead of %d.", actual, expected)
	}
}