/**
 * 203 - Easy
 * <p>
 * Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.
 * <p>
 * Example 1:
 *      Input: head = [1,2,6,3,4,5,6], val = 6
 *      Output: [1,2,3,4,5]
 * <p>
 * Example 2:
 *      Input: head = [], val = 1
 *      Output: []
 * <p>
 * Example 3:
 *      Input: head = [7,7,7,7], val = 7
 *      Output: []
 * <p>
 * Constraints:
 *      - The number of nodes in the list is in the range [0, 10^4].
 *      - 1 <= Node.val <= 50
 *      - 0 <= val <= 50
 */
package removeelements;

import commonclasses.ListNode;

public class RemoveElements {
    public ListNode removeElements(ListNode head, int val) {
        if (head == null) {
            return null;
        }

        while (head != null && head.val == val) {
            head = head.next;
        }

        ListNode current = head;
        ListNode prev = null;

        while (current != null) {
            if (current.val != val) {
                prev = current;
            } else {
                prev.next = current.next;
            }

            current = current.next;
        }

        return head;
    }
}

