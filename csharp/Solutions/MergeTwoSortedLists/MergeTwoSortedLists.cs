/// <summary>
/// 21 - Easy
/// 
/// You are given the heads of two sorted linked lists list1 and list2.
/// Merge the two lists into one sorted list. The list should be made by splicing together
/// the nodes of the first two lists. Return the head of the merged linked list.
/// 
/// Example 1:
///      Input: list1 = [1,2,4], list2 = [1,3,4]
///      Output: [1,1,2,3,4,4]
/// 
/// Example 2:
///      Input: list1 = [], list2 = []
///      Output: []
/// 
/// Example 3:
///      Input: list1 = [], list2 = [0]
///      Output: [0]
/// 
/// Constraints:
///      - The number of nodes in both lists is in the range [0, 50].
///      - -100 <= Node.val <= 100
///      - Both list1 and list2 are sorted in non-decreasing order.
/// </summary>
using Solutions.CommonClases;

namespace Solutions.MergeTwoSortedLists;

public class MergeTwoSortedLists
{
    public ListNode? MergeTwoLists(ListNode? list1, ListNode? list2) {
        ListNode result = new();

        if (list1 == null && list2 == null)
        {
            return null;
        }
        else
        {
            if (list1 == null)
            {
                if (list2 != null) result.val = list2.val;
                result.next = MergeTwoLists(list1, list2?.next);
            }
            else if (list2 == null)
            {
                result.val = list1.val;
                result.next = MergeTwoLists(list1.next, list2);
            }
            else if (list1.val < list2.val)
            {
                result.val = list1.val;
                result.next = MergeTwoLists(list1.next, list2);
            }
            else
            {
                result.val = list2.val;
                result.next = MergeTwoLists(list1, list2.next);
            }
        }

        return result;
    }
}

