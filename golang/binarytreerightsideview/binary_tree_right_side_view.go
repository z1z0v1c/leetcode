/**
 * 199 - Medium
 * Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
 * 
 * Example 1:
 * 
 * 	Input: root = [1,2,3,null,5,null,4]
 * 	Output: [1,3,4]
 * 
 * Example 2:
 * 
 * 	Input: root = [1,2,3,4,null,null,null,5]
 * 	Output: [1,3,4,5]
 * 
 * Example 3:
 * 
 * 	Input: root = [1,null,3]
 * 	Output: [1,3]
 * 
 * Example 4:
 * 
 * 	Input: root = []
 * 	Output: []
 * 
 * Constraints:
 * 
 * 	- The number of nodes in the tree is in the range [0, 100].
 * 	- -100 <= Node.val <= 100
 */
package binarytreerightsideview

import cs "github.com/z1z0v1c/leetcode/commonstructs"

/**
 * >>> Deapth first search <<<
 */
func rightSideView(root *cs.TreeNode) []int {
	if root == nil {
		return nil
	}

	rsw := make([]int, 0)

	rightSideViewDFS(root, &rsw, 0)	

	return rsw
}

func rightSideViewDFS(root *cs.TreeNode, rsw *[]int, level int) {
	if root == nil {
		return
	}

	if len(*rsw) == level {
		*rsw = append(*rsw, root.Val)
	}

	rightSideViewDFS(root.Right, rsw, level+1)
	rightSideViewDFS(root.Left, rsw, level+1)
}

/**
 * >>> Breat first search <<<
 * 
 * func rightSideView(root *cs.TreeNode) []int {
 * 	if root == nil {
 * 		return nil
 * 	}
 * 
 * 	rsw := make([]int, 0)
 * 	queue := []*cs.TreeNode{root}
 * 
 * 	for len(queue) > 0 {
 * 		length := len(queue)
 * 		rsw = append(rsw, queue[length-1].Val)
 * 		
 * 		for i := 0; i < length; i++ {
 * 			node := queue[0]
 * 			queue = queue[1:]
 * 			
 * 			if node.Left != nil {
 * 				queue = append(queue, node.Left)
 * 			}
 * 
 * 			if node.Right != nil {
 * 				queue = append(queue, node.Right)
 * 			}
 * 		}
 * 	}
 * 
 * 	return rsw
 * }
 */