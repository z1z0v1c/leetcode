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

using Solutions.CommonClasses;

namespace Solutions.RemoveDuplicatesFromList;

public class RemoveDuplicatesFromListSolution
{
    public ListNode? DeleteDuplicates(ListNode? head)
    {
        if (head == null) return head;

        var n = head;
        while (n.Next != null)
        {
            while (n.Next != null && n.Val == n.Next.Val)
            {
                n.Next = n.Next.Next;
            }

            if (n.Next != null) n = n.Next;
        }

        return head;
    }
}