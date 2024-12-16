package removeelements;

import commonclasses.ListNode;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class RemoveElementsTest {
    private RemoveElements removeElements;

    @BeforeEach
    void setUp() {
        removeElements = new RemoveElements();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        int val = 6;
        var head =
                new ListNode(1,
                        new ListNode(2,
                                new ListNode(6,
                                        new ListNode(3,
                                                new ListNode(4,
                                                        new ListNode(5,
                                                                new ListNode(6)))))));

        var expected =
                new ListNode(1,
                        new ListNode(2,
                                new ListNode(3,
                                        new ListNode(4,
                                                new ListNode(5)))));

        var actual = removeElements.removeElements(head, val);

        while (expected.next != null || actual.next != null) {
            assertEquals(expected.val, actual.val);

            expected = expected.next;
            actual = actual.next;
        }
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        int val = 1;
        ListNode head = null;

        assertNull(removeElements.removeElements(head, val));
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        int val = 7;
        var head =
                new ListNode(7,
                        new ListNode(7,
                                new ListNode(7,
                                        new ListNode(7))));

        assertNull(removeElements.removeElements(head, val));
    }

    @Test
    @DisplayName("Example 4.")
    void testExampleFour() {
        int val = 1;
        var head =
                new ListNode(1,
                        new ListNode(2));

        var expected =
                new ListNode(2);

        var actual = removeElements.removeElements(head, val);

        while (expected.next != null || actual.next != null) {
            assertEquals(expected.val, actual.val);

            expected = expected.next;
            actual = actual.next;
        }
    }
}

