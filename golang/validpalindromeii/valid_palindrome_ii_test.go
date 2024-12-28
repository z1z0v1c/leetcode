package validpalindromeii

import "testing"

func TestValidPalindrome(t *testing.T) {
	// Example one
	s := "aba"
	result := validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}

	// Example two
	s = "abca"
	result = validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}

	// Example three
	s = "abc"
	result = validPalindrome(s)

	if result {
		t.Errorf("validPalindrome(%s) should return false.", s)
	}

	// Example four
	s = "eceec"
	result = validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}
}

