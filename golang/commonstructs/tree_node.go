package commonstructs

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node1 *TreeNode) Equals(node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	equals := node1.Left.Equals(node2.Left)
	if !equals {
		return false
	}

	equals = node1.Right.Equals(node2.Right)

	return equals
}

