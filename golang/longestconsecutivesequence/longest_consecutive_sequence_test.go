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

	// Example three
	nums = []int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}

	expected = 5
	actual = longestConsecutive(nums)

	if expected != actual {
		t.Errorf("longestConsecutive(nums) returned %d instead of %d.", actual, expected)
	}
}
