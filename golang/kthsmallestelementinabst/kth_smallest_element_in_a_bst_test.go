package kthsmallestelementinabst

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestKthSmallest(t *testing.T) {
	// Example one
	k := 1
	root := &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 1,
			Right: &cs.TreeNode{
				Val: 2,
			},
		},
		Right: &cs.TreeNode{
			Val: 4,
		},
	}

	expected := 1
	actual := kthSmallest(root, k)

	if expected != actual {
		t.Errorf("kthSmallest(%#v, %d) returned %d instead of %d.", root, k, actual, expected)
	}

	// Example two
	k = 3
	root = &cs.TreeNode{
		Val: 5,
		Left: &cs.TreeNode{
			Val: 3,
			Left: &cs.TreeNode{
				Val: 2,
				Left: &cs.TreeNode{
					Val: 1,
				},
			},
			Right: &cs.TreeNode{
				Val: 4,
			},
		},
		Right: &cs.TreeNode{
			Val: 6,
		},
	}

	expected = 3
	actual = kthSmallest(root, k)

	if expected != actual {
		t.Errorf("kthSmallest(%#v, %d) returned %d instead of %d.", root, k, actual, expected)
	}
}