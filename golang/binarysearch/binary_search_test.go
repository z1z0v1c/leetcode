package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	expected := 4
	actual := search(nums, target)

	if expected != actual {
		t.Errorf("Error: serch(nums) should returned %d instead %d.", actual, expected)
	}

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 2

	expected = -1
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("Error: serch(nums) should returned %d instead %d.", actual, expected)
	}
}
