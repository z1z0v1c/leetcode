package countgoodnodesinbinarytree

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestGoodNodes(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 1,
			Left: &cs.TreeNode{
				Val: 3,
			},
		},
		Right: &cs.TreeNode{
			Val: 4,
			Left: &cs.TreeNode{
				Val: 1,
			},
			Right: &cs.TreeNode{
				Val: 5,
			},
		},
	}

	expected := 4
	actual := goodNodes(root)

	if expected != actual {
		t.Errorf("goodNodes(%#v) returned %d instead of %d.", root, actual, expected)
	}

	// Example one
	root = &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 3,
			Left: &cs.TreeNode{
				Val: 4,
			},
			Right: &cs.TreeNode{
				Val: 2,
			},
		},
	}

	expected = 3
	actual = goodNodes(root)

	if expected != actual {
		t.Errorf("goodNodes(%#v) returned %d instead of %d.", root, actual, expected)
	}

	// Example three
	root = &cs.TreeNode{
		Val: 1,
	}

	expected = 1
	actual = goodNodes(root)

	if expected != actual {
		t.Errorf("goodNodes(%#v) returned %d instead of %d.", root, actual, expected)
	}
}