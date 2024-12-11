package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	// Example one
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)

	if !result {
		t.Errorf("containsDuplicate function should return true.")
	}

	// Example two
	nums = []int{1, 2, 3, 4}
	result = containsDuplicate(nums)

	if result {
		t.Errorf("containsDuplicate function should return false.")
	}

	// Example three
	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result = containsDuplicate(nums)

	if !result {
		t.Errorf("containsDuplicate function should return true.")
	}
}
