package linkedlistcycleii;

import commonclasses.ListNode;

import java.util.HashSet;
import java.util.Set;

/**
 * 141 - Easy
 * <p>
 * Given head, the head of a linked list, determine if the linked list has a cycle in it.
 * There is a cycle in a linked list if there is some node in the list that can be reached again
 * by continuously following the next pointer. Internally, pos is used to denote the index of the
 * node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
 * <p>
 * Return true if there is a cycle in the linked list. Otherwise, return false.
 * <p>
 * Example 1:
 *      Input: head = [3,2,0,-4], pos = 1
 *      Output: true
 *      Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
 * <p>
 * Example 2:
 *      Input: head = [1,2], pos = 0
 *      Output: true
 *      Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
 * <p>
 * Example 3:
 *      Input: head = [1], pos = -1
 *      Output: false
 *      Explanation: There is no cycle in the linked list.
 * <p>
 * Constraints:
 *      - The number of the nodes in the list is in the range [0, 10^4].
 *      - -10^5 <= Node.val <= 10^5
 *      - pos is -1 or a valid index in the linked-list.
 * <p>
 * Follow up: Can you solve it using O(1) (i.e. constant) memory?
 */

public class LinkedListCycleII {
    public ListNode detectCycle(ListNode head) {
        Set<ListNode> nodes = new HashSet<>();

        while (head != null) {
            if (nodes.contains(head)) {
                return head;
            }

            nodes.add(head);
            head = head.next;
        }

        return null;
    }

    // Keep it for now
    public ListNode detectCycleWrong(ListNode head) {
        if (head == null) {
            return null;
        }

        var turtle = head;
        var beforeRabbit = head;
        var rabbit = head.next;

        while (turtle != rabbit) {
            if (rabbit == null || rabbit.next == null) {
                return null;
            }

            turtle = turtle.next;
            beforeRabbit = rabbit.next;
            rabbit = rabbit.next.next;
        }

        return beforeRabbit;
    }
}
