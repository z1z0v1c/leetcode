package longestrepeatingcharreplacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	// Example one
	s, k := "ABAB", 2

	expected := 4
	actual := characterReplacement(s, k)

	if expected != actual {
		t.Errorf("characterReplacement(%s, %d) returned %d instead of %d.", s, k, actual, expected)
	}

	// Example two
	s, k = "AABABBA", 1

	expected = 4
	actual = characterReplacement(s, k)

	if expected != actual {
		t.Errorf("characterReplacement(%s, %d) returned %d instead of %d.", s, k, actual, expected)
	}

	// Example three
	s, k = "ABAA", 0

	expected = 2
	actual = characterReplacement(s, k)

	if expected != actual {
		t.Errorf("characterReplacement(%s, %d) returned %d instead of %d.", s, k, actual, expected)
	}

	// Example four
	s, k = "ABBB", 2

	expected = 4
	actual = characterReplacement(s, k)

	if expected != actual {
		t.Errorf("characterReplacement(%s, %d) returned %d instead of %d.", s, k, actual, expected)
	}	
	
	// Example four
	s, k = "BAAAB", 2

	expected = 5
	actual = characterReplacement(s, k)

	if expected != actual {
		t.Errorf("characterReplacement(%s, %d) returned %d instead of %d.", s, k, actual, expected)
	}
}