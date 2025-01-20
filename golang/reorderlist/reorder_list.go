/**
 * 143 - Medium
 *
 * You are given the head of a singly linked-list. The list can be represented as:
 *
 * 	L0 → L1 → … → Ln - 1 → Ln
 *
 * Reorder the list to be on the following form:
 *
 * 	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 * You may not modify the values in the list's nodes. Only nodes themselves may be changed.
 *
 * Example 1:
 *
 * 	Input: head = [1,2,3,4]
 * 	Output: [1,4,2,3]
 *
 * Example 2:
 *
 * 	Input: head = [1,2,3,4,5]
 * 	Output: [1,5,2,4,3]
 *
 * Constraints:
 *
 * 	- The number of nodes in the list is in the range [1, 5 * 10^4].
 * 	- 1 <= Node.val <= 1000
 */
package reorderlist

import cs "github.com/z1z0v1c/leetcode/commonstructs"

func reorderList(head *cs.ListNode) {
	// head is never nil
	if head.Next == nil {
		return
	}

	slow, fast := head, head.Next

	// Find the middle node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	half := make([]*cs.ListNode, 0)

	// Get second half of the list
	for slow.Next != nil {
		half = append(half, slow.Next)
		slow = slow.Next
	}

	node := head

	// Insert nodes
	for i := len(half) - 1; i >= 0; i-- {
		next := node.Next

		node.Next = half[i]
		half[i].Next = next

		node = node.Next.Next
	}

	// Update the last node
	node.Next = nil
}
