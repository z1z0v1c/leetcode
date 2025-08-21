namespace Solutions.CommonClases;

public class TreeNode
{
    public int Val;
    public TreeNode Left;
    public TreeNode Right;

    public TreeNode(int val) {
        this.Val = val;
    }

    public TreeNode(int val, TreeNode left, TreeNode right) {
        this.Val = val;
        this.Left = left;
        this.Right = right;
    }
}