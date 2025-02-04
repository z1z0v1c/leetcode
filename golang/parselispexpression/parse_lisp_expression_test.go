package parselispexpression

import "testing"

func TestEvaluate(t *testing.T) {
	// Example one
	// expression := "(let x 2 (mult x (let x 3 y 4 (add x y))))"

	// expected := 14
	// actual := evaluate(expression)

	// if expected != actual {
	// 	t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	// }

	// Example six
	expression := "(let x 2 y 4 (add x y))"

	expected := 6
	actual := evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example two
	expression = "(let x 3 x 2 x)"

	expected = 2
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example three
	// expression = "(let x 1 y 2 x (add x y) (add x y))"

	// expected = 5
	// actual = evaluate(expression)

	// if expected != actual {
	// 	t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	// }

	// Example four
	expression = "7"
	
	expected = 7
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example five
	expression = "123"
	
	expected = 123
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}
}
