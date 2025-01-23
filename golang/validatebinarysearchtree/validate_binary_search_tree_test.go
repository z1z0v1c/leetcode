package validatebinarysearchtree

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestIsValidBST(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 2,
		Left: &cs.TreeNode{
			Val: 1,
		},
		Right: &cs.TreeNode{
			Val: 3,
		},
	}

	if !isValidBST(root) {
		t.Errorf("isValidBST(%#v) shoud return true for example 1.", root)
	}

	// Example two
	root = &cs.TreeNode{
		Val: 5,
		Left: &cs.TreeNode{
			Val: 1,
		},
		Right: &cs.TreeNode{
			Val: 4,
			Left: &cs.TreeNode{
				Val: 3,
			},
			Right: &cs.TreeNode{
				Val: 6,
			},
		},
	}

	if isValidBST(root) {
		t.Errorf("isValidBST(%#v) shoud return false for example 2.", root)
	}

	// Example three
	root = &cs.TreeNode{
		Val: 2,
		Left: &cs.TreeNode{
			Val: 2,
		},
		Right: &cs.TreeNode{
			Val: 2,
		},
	}

	if isValidBST(root) {
		t.Errorf("isValidBST(%#v) shoud return false for example 3.", root)
	}

	// Example four
	root = &cs.TreeNode{
		Val: 45,
		Left: &cs.TreeNode{
			Val: 42,
			Right: &cs.TreeNode{
				Val: 44,
				Left: &cs.TreeNode{
					Val: 43,
					Left: &cs.TreeNode{
						Val: 41,
					},
				},
			},
		},
	}

	if isValidBST(root) {
		t.Errorf("isValidBST(%#v) shoud return false for example 4.", root)
	}
}
