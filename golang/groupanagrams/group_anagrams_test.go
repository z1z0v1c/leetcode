package groupanagrams

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	// Example one
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	actual := groupAnagrams(strs)

	if len(expected) != len(actual) {
		t.Error("groupAnagrams(strs) returned incorrect result for example 1.")
	}
	 
	for i := 0; i < len(expected); i++ {
		if !slices.Equal(expected[i], actual[i]) {
			t.Error("groupAnagrams(strs) returned incorrect result for example 1.")
		}
	}

	// Example two
	strs = []string{""}

	expected = [][]string{{""}}
	actual = groupAnagrams(strs)

	if !slices.Equal(expected[0], actual[0]) {
		t.Error("groupAnagrams(strs) returned incorrect result for example 2.")
	}

	// Example three
	strs = []string{"a"}

	expected = [][]string{{"a"}}
	actual = groupAnagrams(strs)

	if !slices.Equal(expected[0], actual[0]) {
		t.Error("groupAnagrams(strs) returned incorrect result for example 3.")
	}
}
