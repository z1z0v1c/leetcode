package longestpalindromesubstring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := ""
	expected := ""
	result := longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}
}
