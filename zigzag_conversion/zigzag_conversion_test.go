package zigzagconversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	s := ""
	n := 0
	expected := ""
	result := convert(s, n)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}
}
