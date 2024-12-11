package validpalindromeii

import "testing"

func TestValidPalindrome(t *testing.T) {
	s := "aba"
	result := validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}

	s = "abca"
	result = validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}

	s = "abc"
	result = validPalindrome(s)

	if result {
		t.Errorf("validPalindrome(%s) should return false.", s)
	}

	s = "eceec"
	result = validPalindrome(s)

	if !result {
		t.Errorf("validPalindrome(%s) should return true.", s)
	}
}

