package decodeways

import "testing"

func TestNumDecodings(t *testing.T) {
	// Example one
	s := "12"

	expected := 2
	actual := numDecodings(s)

	if expected != actual {
		t.Errorf("numDecodings(%s) returned %d instead of %d.", s, actual, expected)
	}

	// Example two
	s = "226"

	expected = 3
	actual = numDecodings(s)

	if expected != actual {
		t.Errorf("numDecodings(%s) returned %d instead of %d.", s, actual, expected)
	}

	// Example three
	s = "06"

	expected = 0
	actual = numDecodings(s)

	if expected != actual {
		t.Errorf("numDecodings(%s) returned %d instead of %d.", s, actual, expected)
	}

	// Example three
	s = "2611055971756562"

	expected = 4
	actual = numDecodings(s)

	if expected != actual {
		t.Errorf("numDecodings(%s) returned %d instead of %d.", s, actual, expected)
	}
}