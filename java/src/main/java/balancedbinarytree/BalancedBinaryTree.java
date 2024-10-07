// Difficulty - Easy
//
// Given a binary tree, determine if it is height-balanced.
//
// Example 1:
//      Input: root = [3,9,20,null,null,15,7]
//      Output: true
//
// Example 2:
//      Input: root = [1,2,2,3,3,null,null,4,4]
//      Output: false
//
// Example 3:
//      Input: root = []
//      Output: true
//
// Constraints:
//      The number of nodes in the tree is in the range [0, 5000].
//      -104 <= Node.val <= 104

package balancedbinarytree;

import commonclasses.TreeNode;

public class BalancedBinaryTree {
    public boolean isBalanced(TreeNode root) {
        if (root == null) {
            return true;
        }

        int leftHeight = getHeight(root.left);
        int rightHeight = getHeight(root.right);

        boolean isLeftBalanced = isBalanced(root.left);
        if(!isLeftBalanced) {
            return false;
        }

        boolean isRightBalanced = isBalanced(root.right);
        if (!isRightBalanced) {
            return false;
        }

        return Math.abs(leftHeight - rightHeight) <= 1;
    }

    private static int getHeight(TreeNode node) {
        if (node == null) {
            return 0;
        }

        int leftHeight = getHeight(node.left);
        int rightHeight = getHeight(node.right);

        return Math.max(leftHeight, rightHeight) + 1;
    }
}
