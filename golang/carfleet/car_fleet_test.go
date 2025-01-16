package carfleet

import "testing"

func TestCarFleet(t *testing.T) {
	// Example one
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}

	expected := 3
	actual := carFleet(target, position, speed)

	if expected != actual {
		t.Errorf("carFleet(%d, %#v, %#v) returned %d instead of %d.", target, position, speed, actual, expected)
	}

	// Example two
	target = 10
	position = []int{3}
	speed = []int{3}

	expected = 1
	actual = carFleet(target, position, speed)

	if expected != actual {
		t.Errorf("carFleet(%d, %#v, %#v) returned %d instead of %d.", target, position, speed, actual, expected)
	}

	// Example three
	target = 100
	position = []int{0, 2, 4}
	speed = []int{4, 2, 1}

	expected = 1
	actual = carFleet(target, position, speed)

	if expected != actual {
		t.Errorf("carFleet(%d, %#v, %#v) returned %d instead of %d.", target, position, speed, actual, expected)
	}

	// Example four
	target = 10
	position = []int{6, 8}
	speed = []int{3, 2}

	expected = 2
	actual = carFleet(target, position, speed)

	if expected != actual {
		t.Errorf("carFleet(%d, %#v, %#v) returned %d instead of %d.", target, position, speed, actual, expected)
	}

	// Example five
	target = 10
	position = []int{0, 4, 2}
	speed = []int{2, 1, 3}

	expected = 1
	actual = carFleet(target, position, speed)

	if expected != actual {
		t.Errorf("carFleet(%d, %#v, %#v) returned %d instead of %d.", target, position, speed, actual, expected)
	}
}
