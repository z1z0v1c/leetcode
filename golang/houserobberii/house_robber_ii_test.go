package houserobberii

import "testing"

func TestRob(t *testing.T) {
	// Example one
	nums := []int{2, 3, 2}

	expected := 3
	actual := rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example two
	nums = []int{1, 2, 3, 1}

	expected = 4
	actual = rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example three
	nums = []int{1, 2, 3}

	expected = 3
	actual = rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example four
	nums = []int{0}

	expected = 0
	actual = rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example five
	nums = []int{200, 3, 140, 20, 10}

	expected = 340
	actual = rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}
}
