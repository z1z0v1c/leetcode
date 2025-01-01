package diameterofbinarytree

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 1,
		Left: &cs.TreeNode{
			Val: 2,
			Left: &cs.TreeNode{
				Val: 4,
			},
			Right: &cs.TreeNode{
				Val: 5,
			},
		},
		Right: &cs.TreeNode{
			Val: 3,
		},
	}	

	expected := 3
	actual := diameterOfBinaryTree(root)

	if expected != actual {
		t.Errorf("diameterOfBinaryTree(root) returned %d instead of %d.", actual, expected)
	}

	// Example one
	root = &cs.TreeNode{
		Val: 1,
		Left: &cs.TreeNode{
			Val: 2,
		},
	}	

	expected = 1
	actual = diameterOfBinaryTree(root)

	if expected != actual {
		t.Errorf("diameterOfBinaryTree(root) returned %d instead of %d.", actual, expected)
	}
}