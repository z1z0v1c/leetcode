/**
 * 105 - Medium
 *
 * Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and
 * inorder is the inorder traversal of the same tree, construct and return the binary tree.
 *
 * Example 1:
 *
 * 	Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 	Output: [3,9,20,null,null,15,7]
 *
 * Example 2:
 *
 * 	Input: preorder = [-1], inorder = [-1]
 * 	Output: [-1]
 *
 * Constraints:
 *
 * 	- 1 <= preorder.length <= 3000
 * 	- inorder.length == preorder.length
 * 	- -3000 <= preorder[i], inorder[i] <= 3000
 * 	- preorder and inorder consist of unique values.
 * 	- Each value of inorder also appears in preorder.
 * 	- preorder is guaranteed to be the preorder traversal of the tree.
 * 	- inorder is guaranteed to be the inorder traversal of the tree.
 */
package constructbinarytreeprein

import (
	"slices"

	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

func buildTree(preorder []int, inorder []int) *cs.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &cs.TreeNode{Val: preorder[0]}

	i := slices.Index(inorder, preorder[0])

	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
