package lowestcommonancestorofbst

import (
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestLowestCommonAncestor(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 6,
		Left: &cs.TreeNode{
			Val: 2,
			Left: &cs.TreeNode{
				Val: 0,
			},
			Right: &cs.TreeNode{
				Val: 4,
				Left: &cs.TreeNode{
					Val: 3,
				},
				Right: &cs.TreeNode{
					Val: 5,
				},
			},
		},
		Right: &cs.TreeNode{
			Val: 8,
			Left: &cs.TreeNode{
				Val: 7,
			},
			Right: &cs.TreeNode{
				Val: 9,
			},
		},
	}

	p := &cs.TreeNode{Val: 2}
	q := &cs.TreeNode{Val: 8}

	expected := 6
	actual := lowestCommonAncestor(root, p, q).Val

	if expected != actual {
		t.Errorf("lowestCommonAncestor(%#v, %#v, %#v) returned %d instead %d.", root, p, q, actual, expected)
	}

	// Example two
	root = &cs.TreeNode{
		Val: 6,
		Left: &cs.TreeNode{
			Val: 2,
			Left: &cs.TreeNode{
				Val: 0,
			},
			Right: &cs.TreeNode{
				Val: 4,
				Left: &cs.TreeNode{
					Val: 3,
				},
				Right: &cs.TreeNode{
					Val: 5,
				},
			},
		},
		Right: &cs.TreeNode{
			Val: 8,
			Left: &cs.TreeNode{
				Val: 7,
			},
			Right: &cs.TreeNode{
				Val: 9,
			},
		},
	}

	p = &cs.TreeNode{Val: 2}
	q = &cs.TreeNode{Val: 4}

	expected = 2
	actual = lowestCommonAncestor(root, p, q).Val

	if expected != actual {
		t.Errorf("lowestCommonAncestor(%#v, %#v, %#v) returned %d instead %d.", root, p, q, actual, expected)
	}

	// Example three
	root = &cs.TreeNode{
		Val: 2,
		Left: &cs.TreeNode{
			Val: 1,
		},
	}

	p = &cs.TreeNode{Val: 2}
	q = &cs.TreeNode{Val: 1}

	expected = 2
	actual = lowestCommonAncestor(root, p, q).Val

	if expected != actual {
		t.Errorf("lowestCommonAncestor(%#v, %#v, %#v) returned %d instead %d.", root, p, q, actual, expected)
	}
}