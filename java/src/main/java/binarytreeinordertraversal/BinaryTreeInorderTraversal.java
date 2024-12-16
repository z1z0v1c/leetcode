/**
 * 94 - Easy
 * <p>
 * Given the root of a binary tree, return the inorder traversal of its nodes' values.
 * <p>
 * Example 1:
 *      Input: root = [1,null,2,3]
 *      Output: [1,3,2]
 * <p>
 * Example 2:
 *      Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
 *      Output: [4,2,6,5,7,1,3,9,8]
 * <p>
 * Example 3:
 *      Input: root = []
 *      Output: []
 * <p>
 * Example 4:
 *      Input: root = [1]
 *      Output: [1]
 * <p>
 * Constraints:
 *      - The number of nodes in the tree is in the range [0, 100].
 *      - -100 <= Node.val <= 100
 * <p>
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */
package binarytreeinordertraversal;

import commonclasses.TreeNode;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class BinaryTreeInorderTraversal {
    public List<Integer> inorderTraversal(TreeNode root) {
        var values = new ArrayList<Integer>();
        var subtrees = new Stack<TreeNode>();

        var current = root;
        while (current != null || !subtrees.isEmpty()) {
            while (current != null) {
                subtrees.push(current);
                current = current.left;
            }

            current = subtrees.pop();
            values.add(current.val);

            current = current.right;
        }

        return values;
    }
}

