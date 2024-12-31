package invertbinarytree

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestInvertTree(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 4,
		Left: &cs.TreeNode{
			Val: 2,
			Left: &cs.TreeNode{
				Val: 1,
			},
			Right: &cs.TreeNode{
				Val: 3,
			},
		},
		Right: &cs.TreeNode{
			Val: 7,
			Left: &cs.TreeNode{
				Val: 6,
			},
			Right: &cs.TreeNode{
				Val: 9,
			},
		},
	}
		
	expected := &cs.TreeNode{
		Val: 4,
		Left: &cs.TreeNode{
			Val: 7,
			Right: &cs.TreeNode{
				Val: 6,
			},
			Left: &cs.TreeNode{
				Val: 9,
			},
		},
		Right: &cs.TreeNode{
			Val: 2,
			Right: &cs.TreeNode{
				Val: 1,
			},
			Left: &cs.TreeNode{
				Val: 3,
			},
		},
	}

	actual := invertTree(root)

	if !expected.Equals(actual) {
		t.Error("invertTree() func returned incorrect value")
	}

	// Example two
	root = &cs.TreeNode{
		Val: 2,
		Left: &cs.TreeNode{
			Val: 1,
		},
		Right: &cs.TreeNode{
			Val: 3,
		},
	}
		
	expected = &cs.TreeNode{
		Val: 2,
		Left: &cs.TreeNode{
			Val: 3,
		},
		Right: &cs.TreeNode{
			Val: 1,
		},
	}

	actual = invertTree(root)

	if !expected.Equals(actual) {
		t.Error("invertTree() func returned incorrect value")
	}

	// Example three
	if invertTree(nil) != nil {
		t.Error("invertTree() func returned incorrect value")
	}
}