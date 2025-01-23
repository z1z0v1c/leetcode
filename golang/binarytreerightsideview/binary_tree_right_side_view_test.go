package binarytreerightsideview

import (
	"slices"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestRightSideView(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 1,
		Left: &cs.TreeNode{
			Val: 2,
			Right: &cs.TreeNode{
				Val: 5,
			},
		},
		Right: &cs.TreeNode{
			Val: 3,
			Right: &cs.TreeNode{
				Val: 4,
			},
		},
	}

	expected := []int{1, 3, 4}
	actual := rightSideView(root)

	if !slices.Equal(expected, actual) {
		t.Errorf("rightSideView(%#v) returned %#v instead of %#v.", root, actual, expected)
	}

	// Example two
	root = &cs.TreeNode{
		Val: 1,
		Left: &cs.TreeNode{
			Val: 2,
			Left: &cs.TreeNode{
				Val: 4,
				Left: &cs.TreeNode{
					Val: 5,
				},
			},
		},
		Right: &cs.TreeNode{
			Val: 3,
		},
	}

	expected = []int{1, 3, 4, 5}
	actual = rightSideView(root)

	if !slices.Equal(expected, actual) {
		t.Errorf("rightSideView(%#v) returned %#v instead of %#v.", root, actual, expected)
	}

	// Example three
	root = &cs.TreeNode{
		Val: 1,
		Right: &cs.TreeNode{
			Val: 3,
		},
	}

	expected = []int{1, 3}
	actual = rightSideView(root)

	if !slices.Equal(expected, actual) {
		t.Errorf("rightSideView(%#v) returned %#v instead of %#v.", root, actual, expected)
	}

	// Example four
	if rightSideView(nil) != nil {
		t.Errorf("rightSideView(%#v) returned %#v instead of %#v.", nil, nil, nil)
	}
}