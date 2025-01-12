package productofarrayexceptself

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	// Example one
	nums := []int{1, 2, 3, 4}

	expected := []int{24, 12, 8, 6}
	actual := productExceptSelf(nums)

	if !slices.Equal(expected, actual) {
		t.Error("productExceptSelf(nums) returned incorrect result in exawmple 1.")
	}

	// Example two
	nums = []int{-1, 1, 0, -3, 3}

	expected = []int{0, 0, 9, 0, 0}
	actual = productExceptSelf(nums)

	if !slices.Equal(expected, actual) {
		t.Error("productExceptSelf(nums) returned incorrect result in exawmple 2.")
	}
}