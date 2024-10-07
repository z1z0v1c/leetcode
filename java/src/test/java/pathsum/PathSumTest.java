package pathsum;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class PathSumTest {
    private PathSum pathSum;

    @BeforeEach
    void setUp() {
        pathSum = new PathSum();
    }

    @Test
    @DisplayName("Example one")
    void testIsBalancedExampleOne() {
        int targetSum = 22;
        var root = new TreeNode(
                5,
                new TreeNode(
                        4,
                        new TreeNode(
                                11,
                                new TreeNode(7),
                                new TreeNode(2)
                        ),
                        null

                ),
                new TreeNode(
                        8,
                        new TreeNode(13),
                        new TreeNode(
                                4,
                                null,
                                new TreeNode(1)
                        )
                )
        );

        assertTrue(pathSum.hasPathSum(root, targetSum), "Should return true");
    }

    @Test
    @DisplayName("Example two")
    void testIsBalancedExampleTwo() {
        int targetSum = 5;
        var root = new TreeNode(
                1,
                new TreeNode(2),
                new TreeNode(3)
        );

        assertFalse(pathSum.hasPathSum(root, targetSum), "Should return true");
    }

    @Test
    @DisplayName("Example three")
    void testIsBalancedExampleThree() {
        int targetSum = 0;
        TreeNode root = null;

        assertFalse(pathSum.hasPathSum(root, targetSum), "Should return true");
    }
}
