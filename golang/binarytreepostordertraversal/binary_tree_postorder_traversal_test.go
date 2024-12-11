package binarytreepostordertraversal

import (
	"slices"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestPostorderTraversal(t *testing.T) {
	tn0 := &cs.TreeNode{Val: 1}
	tn1 := &cs.TreeNode{Val: 2}
	tn2 := &cs.TreeNode{Val: 3}

	tn0.Right = tn1
	tn1.Left = tn2

	expected := []int{3, 2, 1}
	actual := postorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("postorderTraversal() returned incorrect value.")
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
	tn4.Right = tn6
	tn2.Right = tn7
	tn7.Left = tn8

	expected = []int{4, 6, 7, 5, 2, 9, 8, 3, 1}
	actual = postorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("postorderTraversal() returned incorrect value.")
	}

	tn0 = &cs.TreeNode{Val: 1}

	expected = []int{1}
	actual = postorderTraversal(tn0)

	if !slices.Equal(actual, expected) {
		t.Errorf("postorderTraversal() returned incorrect value.")
	}

	// Test nil input
	if postorderTraversal(nil) != nil {
		t.Errorf("postorderTraversal(nil) should return nil.")
	}
}
