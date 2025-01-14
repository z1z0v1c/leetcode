package dailytemperatures

import (
	"slices"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	// Example one
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	actual := dailyTemperatures(temperatures)

	if !slices.Equal(expected, actual) {
		t.Errorf("dailyTemperatures(%#v) returned %#v instead of %#v.", temperatures, actual, expected)
	}

	// Example two
	temperatures = []int{30, 40, 50, 60}

	expected = []int{1, 1, 1, 0}
	actual = dailyTemperatures(temperatures)

	if !slices.Equal(expected, actual) {
		t.Errorf("dailyTemperatures(%#v) returned %#v instead of %#v.", temperatures, actual, expected)
	}

	// Example three
	temperatures = []int{30, 60, 90}

	expected = []int{1, 1, 0}
	actual = dailyTemperatures(temperatures)

	if !slices.Equal(expected, actual) {
		t.Errorf("dailyTemperatures(%#v) returned %#v instead of %#v.", temperatures, actual, expected)
	}
}
