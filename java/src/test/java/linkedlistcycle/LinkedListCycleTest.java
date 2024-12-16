package linkedlistcycle;

import commonclasses.ListNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LinkedListCycleTest {
    private LinkedListCycle linkedListCycle;

    @BeforeEach
    void setUp() {
        linkedListCycle = new LinkedListCycle();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        var head = new ListNode(3);
        var nodeOne = new ListNode(2);
        var nodeTwo = new ListNode(1);
        var nodeThree = new ListNode(-4);

        head.next = nodeOne;
        nodeOne.next = nodeTwo;
        nodeTwo.next = nodeThree;
        nodeThree.next = nodeOne;

        assertTrue(linkedListCycle.hasCycle(head));
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        var head = new ListNode(1);
        var nodeOne = new ListNode(2);

        head.next = nodeOne;
        nodeOne.next = head;

        assertTrue(linkedListCycle.hasCycle(head));
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        var head = new ListNode(1);

        assertFalse(linkedListCycle.hasCycle(head));
    }

    @Test
    @DisplayName("Example 4.")
    void testExampleFour() {
        var head = new ListNode(7);
        var nodeOne = new ListNode(5);
        var nodeTwo = new ListNode(9);
        var nodeThree = new ListNode(-3);
        var nodeFour = new ListNode(1);

        head.next = nodeOne;
        nodeOne.next = nodeTwo;
        nodeTwo.next = nodeThree;
        nodeThree.next = nodeFour;

        assertFalse(linkedListCycle.hasCycle(head));
    }

    @Test
    @DisplayName("Example 5.")
    void testExampleFive() {
        ListNode head = null;

        assertFalse(linkedListCycle.hasCycle(head));
    }
}

