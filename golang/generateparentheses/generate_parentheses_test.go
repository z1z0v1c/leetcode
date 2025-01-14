package generateparentheses

import (
	"slices"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	// Example one
	n := 3

	expected := []string{"((()))","(()())","(())()","()(())","()()()"}
	actual := generateParenthesis(n)

	if !slices.Equal(expected, actual) {
		t.Errorf("generateParentheses(%d) returned %#v instead of %#v.", n, actual, expected)
	}

	// Example two
	n = 1

	expected = []string{"()"}
	actual = generateParenthesis(n)

	if !slices.Equal(expected, actual) {
		t.Errorf("generateParentheses(%d) returned %#v instead of %#v.", n, actual, expected)
	}
}