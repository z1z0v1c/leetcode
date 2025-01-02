/*
572 - Easy

Given the roots of two binary trees root and subRoot,
return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
The tree tree could also be considered as a subtree of itself.

Example 1:

	Input: root = [3,4,5,1,2], subRoot = [4,1,2]
	Output: true
	
Example 2:

	Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
	Output: false

Constraints:

	- The number of nodes in the root tree is in the range [1, 2000].
	- The number of nodes in the subRoot tree is in the range [1, 1000].
	- -10^4 <= root.val <= 10^4
	- -10^4 <= subRoot.val <= 10^4
*/
package subtreeofanothertree

import cs "github.com/z1z0v1c/leetcode/commonstructs"

func isSubtree(root *cs.TreeNode, subRoot *cs.TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	}

	isSame := equals(root, subRoot)
	if isSame {
		return true
	}
	
	left := isSubtree(root.Left, subRoot)
	right := isSubtree(root.Right, subRoot) 

	return left || right
}

func equals(root1 *cs.TreeNode, root2 *cs.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}	

	isSame := equals(root1.Left, root2.Left)
	if !isSame {
		return false
	}

	return equals(root1.Right, root2.Right)
}