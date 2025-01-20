package commonstructs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node1 *ListNode) Equals(node2 *ListNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return node1.Next.Equals(node2.Next)
}
