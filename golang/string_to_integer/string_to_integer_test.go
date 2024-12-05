package stringtointeger

import "testing"

func TestMyAtoi(t *testing.T) {
	s := ""
	expected := 0
	result := myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "     "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "   -  "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "  +   "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "  -=   "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "  =   "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "  -0  "
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "42"
	expected = 42
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = " -042"
	expected = -42
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "1337c0d3"
	expected = 1337
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "0-1"
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "words and 987"
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s = "-91283472332"
	expected = -2147483648
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}

	s ="+-12"
	expected = 0
	result = myAtoi(s)

	if result != expected {
		t.Errorf("reverse(%s) returned %d, expected %d", s, result, expected)
	}
}
