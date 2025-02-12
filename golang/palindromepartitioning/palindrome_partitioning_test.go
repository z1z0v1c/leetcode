package palindromepartitioning

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	// Example one
	s := "aab"

	expected := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	actual := partition(s)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("partition(%s) returned %#v instead of %#v.", s, actual, expected)
	}

	// Example two
	s = "a"

	expected = [][]string{{"a"}}
	actual = partition(s)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("partition(%s) returned %#v instead of %#v.", s, actual, expected)
	}

	// Example two
	s = "cdd"

	expected = [][]string{{"c", "d", "d"}, {"c", "dd"}}
	actual = partition(s)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("partition(%s) returned %#v instead of %#v.", s, actual, expected)
	}
}