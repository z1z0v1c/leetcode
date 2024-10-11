package binarytreeinordertraversal;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class BinaryTreeInorderTraversalTest {
    private BinaryTreeInorderTraversal binaryTreeInorderTraversal;

    @BeforeEach
    void setUp() {
        binaryTreeInorderTraversal = new BinaryTreeInorderTraversal();
    }

    @Test
    @DisplayName("Example one")
    void testExampleOne() {
        var root = new TreeNode(
                1,
                null,
                new TreeNode(
                        2,
                        new TreeNode(3 ),
                        null
                )
        );

        var expected = Arrays.asList(1, 3, 2);
        var actual = binaryTreeInorderTraversal.inorderTraversal(root);

        for (int i = 0; i < expected.size(); i ++) {
            assertEquals(expected.get(i), actual.get(i) , "Elements should be equals");
        }
    }

    @Test
    @DisplayName("Example two")
    void testExampleTwo() {
        var root = new TreeNode(
                1,
                new TreeNode(
                        2,
                        new TreeNode(4),
                        new TreeNode(
                                5,
                                new TreeNode(6),
                                new TreeNode(7)
                        )
                ),
                new TreeNode(
                        3,
                        null,
                        new TreeNode(
                                8,
                                new TreeNode(9),
                                null
                        )
                )
        );

        var expected = Arrays.asList(4, 2, 6, 5, 7, 1, 3, 9, 8);
        var actual = binaryTreeInorderTraversal.inorderTraversal(root);

        for (int i = 0; i < expected.size(); i ++) {
            assertEquals(expected.get(i), actual.get(i) , "Elements should be equals.");
        }
    }

    @Test
    @DisplayName("Example three")
    void testExampleThree() {
        TreeNode root = null;

        var actual = binaryTreeInorderTraversal.inorderTraversal(root);

        assertTrue(actual.isEmpty(), "The list should be empty.");
    }

    @Test
    @DisplayName("Example four")
    void testExampleFour() {
        var root = new TreeNode(
                1,
                null,
                new TreeNode(
                        2,
                        new TreeNode(3 ),
                        null
                )
        );

        var expected = Arrays.asList(1, 3, 2);
        var actual = binaryTreeInorderTraversal.inorderTraversal(root);

        for (int i = 0; i < expected.size(); i ++) {
            assertEquals(expected.get(i), actual.get(i) , "Elements should be equals");
        }
    }
}
