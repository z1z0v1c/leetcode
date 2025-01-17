package kokoeatingbananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	// Example one
	piles, h := []int{3, 6, 7, 11}, 8

	expected := 4
	actual := minEatingSpeed(piles, h)

	if expected != actual {
		t.Errorf("minEatingSpeed(%#v, %d) returned %d instead of %d.", piles, h, actual, expected)
	}

	// Example one
	piles, h = []int{30, 11, 23, 4, 20}, 5

	expected = 30 
	actual = minEatingSpeed(piles, h)

	if expected != actual {
		t.Errorf("minEatingSpeed(%#v, %d) returned %d instead of %d.", piles, h, actual, expected)
	}
	
	// Example one
	piles, h = []int{30, 11, 23, 4, 20}, 6

	expected = 23 
	actual = minEatingSpeed(piles, h)

	if expected != actual {
		t.Errorf("minEatingSpeed(%#v, %d) returned %d instead of %d.", piles, h, actual, expected)
	}
}