package binarytreelevelordertraversal

import (
	"reflect"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestLevelOrder(t *testing.T) {
	// Example one
	root := &cs.TreeNode{
		Val: 3,
		Left: &cs.TreeNode{
			Val: 9,
		},
		Right: &cs.TreeNode{
			Val: 20,
			Left: &cs.TreeNode{
				Val: 15,
			},
			Right: &cs.TreeNode{
				Val: 7,
			},
		},
	}

	expected := [][]int{{3}, {9, 20}, {15, 7}}
	actual := levelOrder(root)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("levelOrder(%#v) returned %#v instead of %3v.", root, actual, expected)
	}

	// Example two
	root = &cs.TreeNode{
		Val: 1,
	}

	expected = [][]int{{1}}
	actual = levelOrder(root)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("levelOrder(%#v) returned %#v instead of %3v.", root, actual, expected)
	}
	
	// Example three
	if levelOrder(nil) != nil {
		t.Errorf("levelOrder(%#v) returned %#v instead of %3v.", nil, actual, nil)
	}
}