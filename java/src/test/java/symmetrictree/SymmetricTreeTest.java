package symmetrictree;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SymmetricTreeTest {
    private SymmetricTree symmetricTree;

    @BeforeEach
    void setUp() {
        symmetricTree = new SymmetricTree();
    }

    @Test
    @DisplayName("Example one")
    void testIsBalancedExampleOne() {
        var root = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(3),
                        new TreeNode(4)
                ),
                new TreeNode(
                        2,
                        new TreeNode(4),
                        new TreeNode(3)
                )
        );

        assertTrue(symmetricTree.isSymmetric(root), "Should return true");
    }

    @Test
    @DisplayName("Example two")
    void testIsBalancedExampleTwo() {
        var root = new TreeNode(
                1,
                new TreeNode(
                        2,
                        null,
                        new TreeNode(3)
                ),
                new TreeNode(
                        2,
                        null,
                        new TreeNode(3)
                )
        );

        assertTrue(symmetricTree.isSymmetric(root), "Should return true");
    }
}
