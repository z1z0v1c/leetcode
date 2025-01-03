/**
 * 100 - Easy
 * <p>
 * Given the roots of two binary trees p and q, write a function to check if they are the same or not.
 * Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
 * <p>
 * Example 1:
 *      Input: p = [1,2,3], q = [1,2,3]
 *      Output: true
 * <p>
 * Example 2:
 *      Input: p = [1,2], q = [1,null,2]
 *      Output: false
 * <p>
 * Example 3:
 *      Input: p = [1,2,1], q = [1,1,2]
 *      Output: false
 * <p>
 * Constraints:
 *      - The number of nodes in both trees is in the range [0, 100].
 *      - -104 <= Node.val <= 104
 */
package sametree;

import commonclasses.TreeNode;

public class SameTree {
    public boolean isSameTree(TreeNode p, TreeNode q) {
        if (p == null && q == null) {
            return true;
        }

        if (p == null || q == null) {
            return false;
        }

        if(p.val != q.val) {
            return false;
        }

        boolean isSame = isSameTree(p.left, q.left);
        if (!isSame) {
            return false;
        }
        isSame = isSameTree(p.right, q.right);

        return isSame;
    }
}

