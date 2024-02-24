package zigzagconversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	s := "A"
	numRows := 1
	expected := "A"
	result := convert(s, numRows)

	if result != expected {
		t.Errorf("convert(%s, %d) returned %s, expected %s", s, numRows, result, expected)
	}

	s = "AHA"
	numRows = 2
	expected = "AAH"
	result = convert(s, numRows)

	if result != expected {
		t.Errorf("convert(%s, %d) returned %s, expected %s", s, numRows, result, expected)
	}
}
