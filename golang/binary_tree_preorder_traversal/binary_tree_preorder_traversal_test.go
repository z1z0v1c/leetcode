package binarytreepreordertraversal

import (
	"slices"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestPreorderTraversal(t *testing.T) {
	tn0 := &cs.TreeNode{Val: 1}
	tn1 := &cs.TreeNode{Val: 2}
	tn2 := &cs.TreeNode{Val: 3}

	tn0.Right = tn1
	tn1.Left = tn2

	expected := []int{1, 2, 3}
	actual := preorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("preorderTraversal() retund incorrect value.")
	}

	tn0 = &cs.TreeNode{Val: 1}
	tn1 = &cs.TreeNode{Val: 2}
	tn2 = &cs.TreeNode{Val: 3}
	tn3 := &cs.TreeNode{Val: 4}
	tn4 := &cs.TreeNode{Val: 5}
	tn5 := &cs.TreeNode{Val: 6}
	tn6 := &cs.TreeNode{Val: 7}
	tn7 := &cs.TreeNode{Val: 8}
	tn8 := &cs.TreeNode{Val: 9}

	tn0.Left = tn1
	tn0.Right = tn2
	tn1.Left = tn3
	tn1.Right = tn4
	tn4.Left = tn5
	tn5.Right = tn6
	tn2.Left = tn7
	tn2.Right = tn8

	expected = []int{1, 2, 4, 5, 6, 7, 3, 8, 9}
	actual = preorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("preorderTraversal() retund incorrect value.")
	}

	tn0 = &cs.TreeNode{Val: 1}

	expected = []int{1}
	actual = preorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("preorderTraversal() retund incorrect value.")
	}

	// Test nil input
	if preorderTraversal(nil) != nil {
	 	t.Errorf("preorderTraversal(nil) should return nil.")
	}
}
