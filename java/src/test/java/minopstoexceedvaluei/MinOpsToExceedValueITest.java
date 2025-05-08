package minopstoexceedvaluei;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class MinOpsToExceedValueITest {
    private MinOpsToExceedValueI minOps;

    @BeforeEach
    void setUp() {
        minOps = new MinOpsToExceedValueI();
    }

    @Test
    @DisplayName("Example one")
    void testMinOperationsExampleOne() {
        int k = 10;
        int[] nums = new int[] { 2, 11, 10, 1, 3 };

        int expected = 3;
        int actual = minOps.minOperations(nums, k);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example two")
    void testMinOperationsExampleTwo() {
        int k = 1;
        int[] nums = new int[] { 1, 1, 2, 4, 9 };

        int expected = 0;
        int actual = minOps.minOperations(nums, k);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example three")
    void testMinOperationsExampleThree() {
        int k = 9;
        int[] nums = new int[] { 1, 1, 2, 4, 9 };

        int expected = 4;
        int actual = minOps.minOperations(nums, k);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example four")
    void testMinOperationsExampleFour() {
        int k = 39;
        int[] nums = new int[] { 39, 100, 81, 98, 59, 39, 20, 25 };

        int expected = 2;
        int actual = minOps.minOperations(nums, k);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }
}
