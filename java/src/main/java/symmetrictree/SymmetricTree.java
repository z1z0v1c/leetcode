package symmetrictree;

import commonclasses.TreeNode;

import java.util.Stack;

/**
 * 101 - Easy
 * <p>
 * Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
 * <p>
 * Example 1:
 *      Input: root = [1,2,2,3,4,4,3]
 *      Output: true
 * <p>
 * Example 2:
 *      Input: root = [1,2,2,null,3,null,3]
 *      Output: false
 * <p>
 * Constraints:
 *      - The number of nodes in the tree is in the range [1, 1000].
 *      - -100 <= Node.val <= 100
 * <p>
 * Follow up: Could you solve it both recursively and iteratively?
 */

public class SymmetricTree {
    public boolean isSymmetricRecursive(TreeNode root) {
        boolean recursive = isSymmetricRecursive(root.left, root.right);
        boolean iterative = isSymmetricIterative(root.left, root.right);

        if (iterative != recursive) {
            throw new IllegalStateException("Both methods should return the same solution.");
        }

        return recursive;
    }

    public boolean isSymmetricRecursive(TreeNode left, TreeNode right) {
        if (left == null && right == null) {
            return true;
        }

        if (left == null || right == null) {
            return false;
        }

        if (left.val != right.val) {
            return false;
        }

        boolean symmetric = isSymmetricRecursive(left.left, right.right);
        if (!symmetric) {
            return false;
        }
        symmetric = isSymmetricRecursive(left.right, right.left);

        return symmetric;
    }

    public boolean isSymmetricIterative(TreeNode left, TreeNode right) {
        Stack<TreeNode> leftSubtrees = new Stack<>();
        Stack<TreeNode> rightSubtrees = new Stack<>();

        var currentLeft = left;
        var currentRight = right;

        while (currentLeft != null || currentRight != null || !leftSubtrees.isEmpty()  || !rightSubtrees.isEmpty()) {
            while (currentLeft != null) {
                leftSubtrees.push(currentLeft);
                currentLeft = currentLeft.left;
            }

            while (currentRight != null) {
                rightSubtrees.push(currentRight);
                currentRight = currentRight.right;
            }

            if (leftSubtrees.size() != rightSubtrees.size()) {
                return false;
            }

            currentLeft = leftSubtrees.pop();
            currentRight = rightSubtrees.pop();

            if (currentLeft.val != currentRight.val) {
                return false;
            }

            currentLeft = currentLeft.right;
            currentRight = currentRight.left;
        }

        return true;
    }
}
