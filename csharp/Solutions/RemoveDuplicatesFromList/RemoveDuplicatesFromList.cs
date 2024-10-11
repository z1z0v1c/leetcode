using Solutions.CommonClases;

namespace Solutions.RemoveDuplicatesFromList;

/// <summary>
/// 83 - Easy
/// 
/// Given the head of a sorted linked list, delete all duplicates such that each 
/// element appears only once. Return the linked list sorted as well.
/// 
/// Example 1:
///      Input: head = [1,1,2]
///      Output: [1,2]
/// 
/// Example 2:
///      Input: head = [1,1,2,3,3]
///      Output: [1,2,3]
/// 
/// Constraints:
///      - The number of nodes in the list is in the range [0, 300].
///      - -100 <= Node.val <= 100
///      - The list is guaranteed to be sorted in ascending order.
/// </summary>

public class RemoveDuplicatesFromList
{
    public ListNode? DeleteDuplicates(ListNode? head)
    {
        if (head == null)
        {
            return head;
        }

        var n = head;
        while (n.next != null)
        {
            while (n.next != null && n.val == n.next.val)
            {
                n.next = n.next.next;
            }

            if (n.next != null)
            {
                n = n.next;
            }
        }

        return head;
    }
}