package minrightshiftstosortarray;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

public class MinRightShiftsToSortArrayTest {
    private MinRightShiftsToSortArray minShifts;

    @BeforeEach
    void setUp() {
        minShifts = new MinRightShiftsToSortArray();
    }

    @Test
    @DisplayName("Example one")
    void testMinOperationsExampleOne() {
        List<Integer> nums = Arrays.asList(new Integer[] { 3, 4, 5, 1, 2 });

        int expected = 2;
        int actual = minShifts.minimumRightShifts(nums);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example two")
    void testMinOperationsExampleTwo() {
        List<Integer> nums = Arrays.asList(new Integer[] { 1, 3, 5 });

        int expected = 0;
        int actual = minShifts.minimumRightShifts(nums);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example three")
    void testMinOperationsExampleThree() {
        List<Integer> nums = Arrays.asList(new Integer[] { 2, 1, 4 });

        int expected = -1;
        int actual = minShifts.minimumRightShifts(nums);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }

    @Test
    @DisplayName("Example four")
    void testMinOperationsExampleFour() {
        List<Integer> nums = Arrays.asList(new Integer[] { 92, 84, 37, 19, 16, 85, 20, 79, 25, 89 });

        int expected = -1;
        int actual = minShifts.minimumRightShifts(nums);

        assertEquals(expected, actual, String.format("Should return %d", expected));
    }
}
