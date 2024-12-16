package sametree;

import commonclasses.TreeNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SameTreeTest {
    private SameTree sameTree;

    @BeforeEach
    void setUp() {
        sameTree = new SameTree();
    }

    @Test
    @DisplayName("Example one")
    void testExampleOne() {
        var p = new TreeNode(
                1,
                new TreeNode(2),
                null
        );

        var q = new TreeNode(
                1,
                null,
                new TreeNode(2)
        );

        assertFalse(sameTree.isSameTree(p, q));
    }

    @Test
    @DisplayName("Example two")
    void testExampleTwo() {
        var p = new TreeNode(
                1,
                new TreeNode(2),
                new TreeNode(3)
        );

        var q = new TreeNode(
                1,
                new TreeNode(2),
                new TreeNode(3)
        );

        assertTrue(sameTree.isSameTree(p, q));
    }

    @Test
    @DisplayName("Example three")
    void testExampleThree() {
        var p = new TreeNode(
                1,
                new TreeNode(2),
                null
        );

        var q = new TreeNode(
                1,
                null,
                new TreeNode(2)
        );

        assertFalse(sameTree.isSameTree(p, q));
    }

    @Test
    @DisplayName("Example four")
    void testExampleFour() {
        var p = new TreeNode(
                10,
                new TreeNode(5,
                    null,
                    new TreeNode(15)),
                null
        );

        var q = new TreeNode(
                10,
                new TreeNode(5),
                new TreeNode(2)
        );

        assertFalse(sameTree.isSameTree(p, q));
    }
}

