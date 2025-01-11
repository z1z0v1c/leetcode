package numberof1bits

import "testing"

func TestHammingWeight(t *testing.T) {
	// Example one
	n := 11
	expected := 3
	actual := hammingWeight(n)

	if expected != actual {
		t.Errorf("hammingWeight(%d) returned %d instead of %d.", n, actual, expected)
	}

	// Example two
	n = 128
	expected = 1
	actual = hammingWeight(n)

	if expected != actual {
		t.Errorf("hammingWeight(%d) returned %d instead of %d.", n, actual, expected)
	}

	// Example three
	n = 2147483645 
	expected = 30
	actual = hammingWeight(n)

	if expected != actual {
		t.Errorf("hammingWeight(%d) returned %d instead of %d.", n, actual, expected)
	}
}