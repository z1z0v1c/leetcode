package parselispexpression

import "testing"

func TestEvaluate(t *testing.T) {
	// Example one
	expression := "7"
	
	expected := 7
	actual := evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example two
	expression = "-7"
	
	expected = -7
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example three
	expression = "123"
	
	expected = 123
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example four
	expression = "-123"
	
	expected = -123
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example five
	expression = "(let x 3 x 2 x)"

	expected = 2
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example six
	expression = "(let x 2 y 4 (add x y))"

	expected = 6
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example seven
	expression = "(let x 2 (mult x (let x 3 y 4 (add x y))))"

	expected = 14
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example eight
	expression = "(let x 1 y 2 x (add x y) (add x y))"

	expected = 5
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example nine
	expression = "(let x 2 (add (let x 3 (let x 4 x)) x))"

	expected = 6
	actual = evaluate(expression)

	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example ten
	expression = "(let x -2 y x y)"

	expected = -2
	actual = evaluate(expression)
	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}

	// Example eleven
	expression = "(let x (add 12 -7) (mult x x))"

	expected = 25
	actual = evaluate(expression)
	if expected != actual {
		t.Errorf("evaluate(%s) returned %d instead of %d.", expression, actual, expected)
	}
}
