/// <summary>
/// Definition for a binary tree node.
/// </summary>

namespace Solutions.CommonClasses;

public class TreeNode(int val, TreeNode? left = null, TreeNode? right = null)
{
    public int Val = val;
    public readonly TreeNode? Left = left;
    public readonly TreeNode? Right = right;
}