package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	// Example one
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 0

	expected := 4
	actual := search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}

	// Example two
	nums, target = []int{4, 5, 6, 7, 0, 1, 2}, 3

	expected = -1
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}

	// Example three
	nums, target = []int{1}, 0

	expected = -1
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}

	// Example four
	nums, target = []int{1, 3}, 3

	expected = 1
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}

	// Example five
	nums, target = []int{3, 1}, 3

	expected = 0
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}

	// Example six
	nums, target = []int{5, 1, 2, 3, 4}, 1

	expected = 1
	actual = search(nums, target)

	if expected != actual {
		t.Errorf("search(%#v, %d) returned %d instead of %d.", nums, target, actual, expected)
	}
}
