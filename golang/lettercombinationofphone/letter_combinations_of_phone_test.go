package lettercombinationofphone

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	// Example one
	digits := "23"

	expected := []string{"ad","ae","af","bd","be","bf","cd","ce","cf"}
	actual := letterCombinations(digits)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("letterCombinations(%s) returned %#v instead of %#v.", digits, actual, expected)
	}

	// Example two
	digits = ""

	expected = []string{}
	actual = letterCombinations(digits)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("letterCombinations(%s) returned %#v instead of %#v.", digits, actual, expected)
	}

	// Example one
	digits = "2"

	expected = []string{"a","b","c"}
	actual = letterCombinations(digits)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("letterCombinations(%s) returned %#v instead of %#v.", digits, actual, expected)
	}
}