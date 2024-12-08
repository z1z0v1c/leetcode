/*

144 - Easy

Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:

	Input: root = [1,null,2,3]
	Output: [1,2,3]

Example 2:

	Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
	Output: [1,2,4,5,6,7,3,8,9]

Example 3:

	Input: root = []
	Output: []

Example 4:

	Input: root = [1]
	Output: [1]

Constraints:

	- The number of nodes in the tree is in the range [0, 100].
	- -100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package binarytreepreordertraversal

import (
	cs "github.com/z1z0v1c/leetcode/commonstructs"
	"github.com/zeroflucs-given/generics/collections/stack"
)

func preorderTraversal(root *cs.TreeNode) []int {
	if root == nil {
		return nil
	}

	var values []int
	values = append(values, root.Val)

	nodes := []*cs.TreeNode{}
	nodes = append(nodes, root)

	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]

		if node.Left != nil {
			values = append(values, node.Left.Val)
			nodes = append(nodes, node.Left)
			node.Left = nil
		} else if node.Right != nil {
			values = append(values, node.Right.Val)
			nodes = append(nodes, node.Right)
			node.Right = nil
		} else {
			nodes = nodes[:len(nodes)-1]
		}
	}

	return values
}

// func preorderTraversalStack(root *cs.TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 
// 	var values []int
// 	values = append(values, root.Val)
// 
// 	nodes := stack.NewStack[*cs.TreeNode](100)
// 	nodes.Push(root)
// 
// 	for nodes.Count() != 0 {
// 		_, node := nodes.Peek()
// 
// 		if node.Left != nil {
// 			values = append(values, node.Left.Val)
// 			nodes.Push(node.Left)
// 			node.Left = nil
// 		} else if node.Right != nil {
// 			values = append(values, node.Right.Val)
// 			nodes.Push(node.Right)
// 			node.Right = nil
// 		} else {
// 			nodes.Pop()
// 		}
// 	}
// 
// 	return values
// }
// 
// func preorderTraversalResursive(root *cs.TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 
// 	var values []int
// 
// 	values = append(values, root.Val)
// 	values = append(values, preorderTraversalResursive(root.Left)...)
// 	values = append(values, preorderTraversalResursive(root.Right)...)
// 
// 	return values
// }
