/**
 * 145 - Easy
 * 
 * Given the root of a binary tree, return the postorder traversal of its nodes' values.
 * 
 * Example 1:
 * 
 * 	Input: root = [1,null,2,3]
 * 	Output: [3,2,1]
 * 
 * Example 2:
 * 
 * 	Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
 * 	Output: [4,6,7,5,2,9,8,3,1]
 * 
 * Example 3:
 * 
 * 	Input: root = []
 * 	Output: []
 * 
 * Example 4:
 * 
 * 	Input: root = [1]
 * 	Output: [1]
 * 
 * Constraints:
 * 
 *     - The number of the nodes in the tree is in the range [0, 100].
 *     - -100 <= Node.val <= 100
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */
package binarytreepostordertraversal

import cs "github.com/z1z0v1c/leetcode/commonstructs"

func postorderTraversal(root *cs.TreeNode) []int {
	if root == nil {
		return nil
	}

	values := []int{}
	nodes := []*cs.TreeNode{}
	nodes = append(nodes, root)

	for len(nodes) != 0 {
		peek := nodes[len(nodes)-1]

		if peek.Left != nil {
			nodes = append(nodes, peek.Left)
		} else if peek.Right != nil {
			nodes = append(nodes, peek.Right)
		} else {
			values = append(values, peek.Val)
			nodes = nodes[:len(nodes)-1]

			if len(nodes) != 0 {
				if nodes[len(nodes)-1].Left != nil {
					nodes[len(nodes)-1].Left = nil
				} else if nodes[len(nodes)-1].Right != nil {
					nodes[len(nodes)-1].Right = nil
				}
			}
		}
	}

	return values
}

