package palindromicsubstrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	// Example one
	s := "abc"

	expected := 3
	actual := countSubstrings(s)

	if expected != actual {
		t.Errorf("countSubstrings(%s) returned %d instead of %d.", s, actual, expected)
	}
	
	// Example two
	s = "aaa"

	expected = 6
	actual = countSubstrings(s)

	if expected != actual {
		t.Errorf("countSubstrings(%s) returned %d instead of %d.", s, actual, expected)
	}

	// Example three
	s = "xkjkqlajprjwefilxgpdpebieswu"

	expected = 30
	actual = countSubstrings(s)

	if expected != actual {
		t.Errorf("countSubstrings(%s) returned %d instead of %d.", s, actual, expected)
	}

	// Example four
	s = "dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg"

	expected = 77
	actual = countSubstrings(s)

	if expected != actual {
		t.Errorf("countSubstrings(%s) returned %d instead of %d.", s, actual, expected)
	}
}