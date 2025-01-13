package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	// Example one
	nums := []int{100, 4, 200, 1, 3, 2}
	
	expected := 4
	actual := longestConsecutive(nums)

	if expected != actual {
		t.Errorf("longestConsecutive(nums) returned %d instead of %d.", actual, expected)
	}

	// Example two
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	
	expected = 9
	actual = longestConsecutive(nums)

	if expected != actual {
		t.Errorf("longestConsecutive(nums) returned %d instead of %d.", actual, expected)
	}
}