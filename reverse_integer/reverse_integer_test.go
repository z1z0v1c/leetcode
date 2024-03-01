package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	x := 0
	expected := 0
	result := reverse(x)

	if result != expected {
		t.Errorf("reverse(%d) returned %d, expected %d", x, result, expected)
	}

	x = 123
	expected = 321
	result = reverse(x)

	if result != expected {
		t.Errorf("reverse(%d) returned %d, expected %d", x, result, expected)
	}
}