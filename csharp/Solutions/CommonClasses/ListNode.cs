/// <summary>
/// Definition for singly-linked list.
/// </summary>
namespace Solutions.CommonClasses;

public class ListNode(int val = 0, ListNode? next = null)
{
    public int Val = val;
    public ListNode? Next = next;
}

