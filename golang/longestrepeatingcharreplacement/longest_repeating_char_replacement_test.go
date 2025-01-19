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
}