package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	// Example one
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	actual := groupAnagrams(strs)

	if len(expected) != len(actual) {
		t.Error("groupAnagrams(strs) returned incorrect result for example 1.")
	}
	 
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("groupAnagrams(%v) returned %v instead of %v.", strs, actual, expected)
	}

	// Example two
	strs = []string{""}

	expected = [][]string{{""}}
	actual = groupAnagrams(strs)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("groupAnagrams(%v) returned %v instead of %v.", strs, actual, expected)
	}

	// Example three
	strs = []string{"a"}

	expected = [][]string{{"a"}}
	actual = groupAnagrams(strs)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("groupAnagrams(%v) returned %v instead of %v.", strs, actual, expected)
	}
}
