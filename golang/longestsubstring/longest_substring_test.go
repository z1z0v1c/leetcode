package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	err := testLengthOfLongestSubstring(" ", 1)
	if err != nil {
		t.Error(err.Error())
	}

	err = testLengthOfLongestSubstring("abcabcbb", 3)
	if err != nil {
		t.Error(err.Error())
	}

	err = testLengthOfLongestSubstring("bbbbb", 1)
	if err != nil {
		t.Error(err.Error())
	}

	err = testLengthOfLongestSubstring("pwwkew", 3)
	if err != nil {
		t.Error(err.Error())
	}
}

func testLengthOfLongestSubstring(s string, expected int) error {
	result := lengthOfLongestSubstring(s)

	if result != expected {
		return fmt.Errorf("lengthOfLongestSubstring(%s) returned %d, expected %d", s, result, expected)
	}

	return nil
}

