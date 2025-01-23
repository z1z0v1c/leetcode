/**
 * 102 - Medium
 * Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
 *
 * Example 1:
 *
 * 	Input: root = [3,9,20,null,null,15,7]
 * 	Output: [[3],[9,20],[15,7]]
 *
 * Example 2:
 *
 * 	Input: root = [1]
 * 	Output: [[1]]
 *
 * Example 3:
 *
 * 	Input: root = []
 * 	Output: []
 *
 * Constraints:
 *
 * 	- The number of nodes in the tree is in the range [0, 2000].
 * 	- -1000 <= Node.val <= 1000
 */
package binarytreelevelordertraversal

import cs "github.com/z1z0v1c/leetcode/commonstructs"

func levelOrder(root *cs.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levels := make([][]int, 0)
	queue := []*cs.TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		level := make([]int, 0, length)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			
			level = append(level, node.Val)
		
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levels = append(levels, level)
	}
	
	return levels
}
