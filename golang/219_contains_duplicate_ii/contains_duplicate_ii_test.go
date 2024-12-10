package containsduplicateii

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	// Example one
	nums, k := []int{1, 2, 3, 1}, 3
	result := containsNearbyDuplicate(nums, k)

	if !result {
		t.Errorf("containsNearbyDuplicate fuction should return true.")
	}

	// Example two
	nums, k = []int{1, 0, 1, 1}, 1
	result = containsNearbyDuplicate(nums, k)

	if !result {
		t.Errorf("containsNearbyDuplicate fuction should return true.")
	}

	// Example three
	nums, k = []int{1, 2, 3, 1, 2, 3}, 2
	result = containsNearbyDuplicate(nums, k)

	if result {
		t.Errorf("containsNearbyDuplicate fuction should return false.")
	}
}
