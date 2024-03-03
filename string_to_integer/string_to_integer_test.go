package stringtointeger

import "testing"

func TestMyAtoi(t *testing.T) {
	s := ""
	expected := 0
	result := myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}
}
