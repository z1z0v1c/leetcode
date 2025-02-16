package houserobber

import "testing"

func TestRob(t *testing.T) {
	// Example one
	nums := []int{1, 2, 3, 1}

	expected := 4
	actual := Rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example two
	nums = []int{2, 7, 9, 3, 1}

	expected = 12
	actual = Rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example three
	nums = []int{0}

	expected = 0
	actual = Rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example four
	nums = []int{2, 1, 1, 2}

	expected = 4
	actual = Rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}

	// Example five
	nums = []int{
        228, 67, 195, 15, 0, 90, 151, 210, 17, 196, 0, 10, 28, 110, 169, 94, 9, 23, 52, 63,
        136, 122, 22, 191, 144, 22, 173, 106, 88, 59, 200, 156, 39, 109, 244, 231, 183, 99,
        114, 15, 114, 32, 57, 36, 117, 151, 177, 106, 229, 188, 178, 47, 96, 191, 25, 180, 153,
        187, 136, 146, 117, 57, 77, 110, 215, 115, 84, 218, 59, 121, 202, 109, 205, 95, 214,
        100, 175, 50, 223, 11, 14, 164, 224, 15, 100, 241, 61, 64, 197, 206, 3, 149, 108, 186,
    }

	expected = 6762
	actual = Rob(nums)

	if expected != actual {
		t.Errorf("rob(%#v) returned %d instead of %d.", nums, actual, expected)
	}
}
