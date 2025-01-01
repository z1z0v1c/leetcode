package subtreeofanothertree

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestIsSubtree(t *testing.T) {
	root := &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 4,
			Left: &cs.TreeNode{
				Val: 1,
			},
			Right: &cs.TreeNode{
				Val: 2,
			},
		},
		Right: &cs.TreeNode{
			Val: 5,
		},
	}

	subRoot := &cs.TreeNode{
		Val: 4,
		Left: &cs.TreeNode{
			Val: 1,
		},
		Right: &cs.TreeNode{
			Val: 2,
		},
	}

	if !isSubtree(root, subRoot) {
		t.Error("isSubtree(root, subRoot) should return true.")
	}

	root = &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 4,
			Left: &cs.TreeNode{
				Val: 1,
			},
			Right: &cs.TreeNode{
				Val: 2,
				Left: &cs.TreeNode{
					Val: 0,
				},
			},
		},
		Right: &cs.TreeNode{
			Val: 5,
		},
	}

	subRoot = &cs.TreeNode{
		Val: 4,
		Left: &cs.TreeNode{
			Val: 1,
		},
		Right: &cs.TreeNode{
			Val: 2,
		},
	}

	if isSubtree(root, subRoot) {
		t.Error("isSubtree(root, subRoot) should return false.")
	}
}