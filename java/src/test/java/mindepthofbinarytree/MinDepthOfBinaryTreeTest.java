package mindepthofbinarytree;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class MinDepthOfBinaryTreeTest {
    private MinDepthOfBinaryTree minDepthOfBinaryTree;

    @BeforeEach
    void setUp() {
        minDepthOfBinaryTree = new MinDepthOfBinaryTree();
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

        assertEquals(2, minDepthOfBinaryTree.minDepth(root), "Should return 2");
    }

    @Test
    @DisplayName("Example two")
    void testIsBalancedExampleTwo() {
        var root = new TreeNode(
                2,
                null,
                new TreeNode(
                        3,
                        null,
                        new TreeNode(
                                4,
                                null,
                                new TreeNode(
                                        5,
                                        null,
                                        new TreeNode(6)
                                )
                        )
                )
        );

        assertEquals(5, minDepthOfBinaryTree.minDepth(root), "Should return 5");
    }
}
