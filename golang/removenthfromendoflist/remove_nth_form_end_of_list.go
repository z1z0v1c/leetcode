/**
 * 19 - Medium
 *
 * Given the head of a linked list, remove the nth node from the end of the list and return its head.
 *
 * Example 1:
 *
 * 	Input: head = [1,2,3,4,5], n = 2
 * 	Output: [1,2,3,5]
 *
 * Example 2:
 *
 * 	Input: head = [1], n = 1
 * 	Output: []
 *
 * Example 3:
 *
 * 	Input: head = [1,2], n = 1
 * 	Output: [1]
 *
 * Constraints:
 *
 * 	- The number of nodes in the list is sz.
 * 	- 1 <= sz <= 30
 * 	- 0 <= Node.val <= 100
 * 	- 1 <= n <= sz
 *
 * Follow up: Could you do this in one pass?
 */
package removenthfromendoflist

import cs "github.com/z1z0v1c/leetcode/commonstructs"

func removeNthFromEnd(head *cs.ListNode, n int) *cs.ListNode {
	nodes := make([]*cs.ListNode, 0, n)

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	if len(nodes) == 1 {
		return nil
	}

	if len(nodes) == n {
		return nodes[1]
	}
	
	index := len(nodes) - n
	
	if n > 1 {
		nodes[index-1].Next = nodes[index+1]
	} else {
		nodes[index-1].Next = nil 
	}

	return nodes[0]
}
