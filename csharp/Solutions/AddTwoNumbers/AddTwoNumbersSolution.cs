/// <summary>
/// 2 - Medium
/// 
/// You are given two non-empty linked lists representing two non-negative integers. 
/// The digits are stored in reverse order, and each of their nodes contains a single digit. 
/// Add the two numbers and return the sum as a linked list.
/// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
/// 
/// Example 1:
///      Input: l1 = [2,4,3], l2 = [5,6,4]
///      Output: [7,0,8]
///      Explanation: 342 + 465 = 807.
/// 
/// Example 2:
///      Input: l1 = [0], l2 = [0]
///      Output: [0]
/// 
/// Example 3:
///      Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
///      Output: [8,9,9,9,0,0,0,1]
///  
/// Constraints:
///      - The number of nodes in each linked list is in the range [1, 100].
///      - 0 <= Node.val <= 9
///      - It is guaranteed that the list represents a number that does not have leading zeros.
/// </summary>
using Solutions.CommonClasses;

namespace Solutions.AddTwoNumbers;

public class AddTwoNumbersSolution
{
    public ListNode AddTwoNumbers(ListNode? l1, ListNode? l2)
    {
        ListNode? l3 = new();
        var start = l3;

        int reminder = 0;

        while (l1 != null || l2 != null || reminder != 0)
        {
            int sum = (l1?.Val ?? 0) + (l2?.Val ?? 0) + reminder;

            reminder = sum / 10;
            sum = sum % 10;

            if (l3 != null) { l3.Val = sum; };

            l1 = l1?.Next;
            l2 = l2?.Next;

            if (l1 != null || l2 != null || reminder != 0)
            {
                if (l3 != null) { l3.Next = new ListNode(); };
                l3 = l3?.Next;
            }
        }

        return start;
    }
}

