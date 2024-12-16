package balancedbinarytree;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class BalancedBinaryTreeTest {
    private BalancedBinaryTree balancedBinaryTree;

    @BeforeEach
    void setUp() {
        balancedBinaryTree = new BalancedBinaryTree();
    }

    @Test
    @DisplayName("Example one")
    void testIsBalancedExampleOne() {
        var root = new TreeNode(
                3,
                new TreeNode(9),
                new TreeNode(
                        20,
                        new TreeNode(15),
                        new TreeNode(7)
                )
        );

        assertTrue(balancedBinaryTree.isBalanced(root), "Should return true");
    }

    @Test
    @DisplayName("Example two")
    void testIsBalancedExampleTwo() {
        var root = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(
                                3,
                                new TreeNode(4),
                                new TreeNode(4)
                        ),
                        new TreeNode(3)
                ),
                new TreeNode(2)
        );

        assertFalse(balancedBinaryTree.isBalanced(root), "Should return false");
    }

    @Test
    @DisplayName("Example three")
    void testIsBalancedExampleThree() {
        assertTrue(balancedBinaryTree.isBalanced(null), "Should return true");
    }
}

