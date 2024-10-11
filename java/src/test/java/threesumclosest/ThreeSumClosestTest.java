package threesumclosest;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ThreeSumClosestTest {
    private ThreeSumClosest threeSumClosest;

    @BeforeEach
    void setUp() {
        threeSumClosest = new ThreeSumClosest();
    }

    @Test
    @DisplayName("Example one")
    void testExampleOne() {
        int[] nums = {-1, 2, 1, -4};
        int target = 1;

        int expected = 2;
        int actual = threeSumClosest.threeSumClosest(nums, target);

        assertEquals(expected, actual, "Should return 2");
    }

    @Test
    @DisplayName("Example two")
    void testExampleTwo() {
        int[] nums = {0, 0, 0};
        int target = 1;

        int expected = 0;
        int actual = threeSumClosest.threeSumClosest(nums, target);

        assertEquals(expected, actual, "Should return 0");
    }

    @Test
    @DisplayName("Example three")
    void testExampleThree() {
        int[] nums = {4, 0, 5, -5, 3, 3, 0, -4, -5};
        int target = -2;

        int expected = -2;
        int actual = threeSumClosest.threeSumClosest(nums, target);

        assertEquals(expected, actual, "Should return -2");
    }
}