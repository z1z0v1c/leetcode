package searcha2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	// Example one
	matrix, target := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3
	
	if !searchMatrix(matrix, target) {
		t.Errorf("searchMatrix(%#v, %d) should return true.", matrix, target)
	}

	// Example two
	matrix, target = [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 13
	
	if searchMatrix(matrix, target) {
		t.Errorf("searchMatrix(%#v, %d) should return false.", matrix, target)
	}

	// Example three
	matrix, target = [][]int{{1}}, 0
	
	if searchMatrix(matrix, target) {
		t.Errorf("searchMatrix(%#v, %d) should return false.", matrix, target)
	}
}