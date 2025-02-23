package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	// Example one
	coins := []int{1, 2, 5}
	amount := 11

	expected := 3
	actual := coinChange(coins, amount)

	if expected != actual {
		t.Errorf("coinChange(%#v, %d) returned %d instead of %d.", coins, amount, actual, expected)
	}

	// Example two
	coins = []int{2}
	amount = 3

	expected = -1
	actual = coinChange(coins, amount)

	if expected != actual {
		t.Errorf("coinChange(%#v, %d) returned %d instead of %d.", coins, amount, actual, expected)
	}

	// Example three
	coins = []int{1}
	amount = 0

	expected = 0
	actual = coinChange(coins, amount)

	if expected != actual {
		t.Errorf("coinChange(%#v, %d) returned %d instead of %d.", coins, amount, actual, expected)
	}

	// Example four
	coins = []int{2, 5, 10, 1}
	amount = 27

	expected = 4
	actual = coinChange(coins, amount)

	if expected != actual {
		t.Errorf("coinChange(%#v, %d) returned %d instead of %d.", coins, amount, actual, expected)
	}
}
