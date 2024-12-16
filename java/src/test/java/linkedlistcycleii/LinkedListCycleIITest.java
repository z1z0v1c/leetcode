package linkedlistcycleii;

import commonclasses.ListNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LinkedListCycleTest {
    private LinkedListCycleII linkedListCycleII;

    @BeforeEach
    void setUp() {
        linkedListCycleII = new LinkedListCycleII();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        var head = new ListNode(3);
        var nodeOne = new ListNode(2);
        var nodeTwo = new ListNode(0);
        var nodeThree = new ListNode(-4);

        head.next = nodeOne;
        nodeOne.next = nodeTwo;
        nodeTwo.next = nodeThree;
        nodeThree.next = nodeOne;

        assertEquals(nodeOne, linkedListCycleII.detectCycle(head));
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        var head = new ListNode(1);
        var nodeOne = new ListNode(2);

        head.next = nodeOne;
        nodeOne.next = head;

        assertEquals(head, linkedListCycleII.detectCycle(head));
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        var head = new ListNode(1);

        assertNull(linkedListCycleII.detectCycle(head));
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

        assertNull(linkedListCycleII.detectCycle(head));
    }

    @Test
    @DisplayName("Example 5.")
    void testExampleFive() {
        ListNode head = null;

        assertNull(linkedListCycleII.detectCycle(head));
    }

    @Test
    @DisplayName("Example 6.")
    void testExampleSix() {
        var head = new ListNode(4);
        var nodeOne = new ListNode(3);
        var nodeTwo = new ListNode(2);
        var nodeThree = new ListNode(1);

        head.next = nodeOne;
        nodeOne.next = nodeTwo;
        nodeTwo.next = nodeThree;
        nodeThree.next = head;

        assertEquals(head, linkedListCycleII.detectCycle(head));
    }
}

