package constructbinarytreeprein

import (
	"reflect"
	"testing"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func TestBuildTree(t *testing.T) {
	// Example one
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	
	expected := &cs.TreeNode{
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

	actual := buildTree(preorder, inorder)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("buildTree(%#v, %#v) returned %#v instead of %#v.", preorder, inorder, actual, expected)
	}
	
	// Example two
	preorder = []int{-1}
	inorder = []int{-1}
	
	expected = &cs.TreeNode{
		Val: -1,
	}

	actual = buildTree(preorder, inorder)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("buildTree(%#v, %#v) returned %#v instead of %#v.", preorder, inorder, actual, expected)
	}
}