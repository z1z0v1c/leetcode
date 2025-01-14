package evaluatereversepolishnotation

import "testing"

func TestEvalRPN(t *testing.T) {
	// Example one
	tokens := []string{"2","1","+","3","*"}

	expected := 9
	actual := evalRPN(tokens)

	if expected != actual {
		t.Errorf("evalRPN(%#v) returned %d instead of %d.", tokens, actual, expected)
	}

	// Example two
	tokens = []string{"4","13","5","/","+"}

	expected = 6
	actual = evalRPN(tokens)

	if expected != actual {
		t.Errorf("evalRPN(%#v) returned %d instead of %d.", tokens, actual, expected)
	}

	// Example three
	tokens = []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	expected = 22
	actual = evalRPN(tokens)

	if expected != actual {
		t.Errorf("evalRPN(%#v) returned %d instead of %d.", tokens, actual, expected)
	}
}