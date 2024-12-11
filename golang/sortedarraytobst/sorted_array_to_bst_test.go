package sortedarraytobst

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestSortedArrayToBST(t *testing.T) {

	// Example one
	nums := []int{-10, -3, 0, 5, 9}

	tn0 := &cs.TreeNode{Val: 0}
	tn1 := &cs.TreeNode{Val: -3}
	tn2 := &cs.TreeNode{Val: 9}
	tn3 := &cs.TreeNode{Val: -10}
	tn4 := &cs.TreeNode{Val: 5}

	tn0.Left = tn1
	tn0.Right = tn2
	tn1.Left = tn3
	tn2.Left = tn4

	expected := tn0
	actual := sortedArrayToBST(nums)

	// Test first option
	failed := false
	if actual.Equals(expected) {
		failed = true
	}

	tn0 = &cs.TreeNode{Val: 0}
	tn1 = &cs.TreeNode{Val: -10}
	tn2 = &cs.TreeNode{Val: 5}
	tn3 = &cs.TreeNode{Val: -3}
	tn4 = &cs.TreeNode{Val: 9}

	tn0.Left = tn1
	tn0.Right = tn2
	tn1.Right = tn3
	tn2.Right = tn4

	expected = tn0
	actual = sortedArrayToBST(nums)

	// Test second option
	if failed && actual.Equals(expected) {
		t.Errorf("sortedArrayToBST() retund incorrect value.")
	}

	// Example two
	nums = []int{1, 3}

	tn0 = &cs.TreeNode{Val: 3}
	tn1 = &cs.TreeNode{Val: 1}
	tn0.Left = tn1

	expected = tn0
	actual = sortedArrayToBST(nums)

	// Test first option
	failed = false
	if actual.Equals(expected) {
		failed = true
	}

	tn0 = &cs.TreeNode{Val: 1}
	tn1 = &cs.TreeNode{Val: 3}
	tn0.Right = tn1

	expected = tn0
	actual = sortedArrayToBST(nums)

	// Test second option
	if failed && actual.Equals(expected) {
		t.Errorf("sortedArrayToBST() retund incorrect value.")
	}
}
