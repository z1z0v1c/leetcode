package mincostclimbingstairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	// Example one
	cost := []int{10, 25, 20}

	expected := 15
	actual := minCostClimbingStairs(cost)

	if expected != actual {
		t.Errorf("minCostClimbingStairs(cost) returned %d instead %d.", actual, expected)
	}

	// Example two
	cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	expected = 6
	actual = minCostClimbingStairs(cost)

	if expected != actual {
		t.Errorf("minCostClimbingStairs(cost) returned %d instead %d.", actual, expected)
	}
}
