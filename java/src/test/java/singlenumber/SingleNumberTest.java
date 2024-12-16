package singlenumber;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SingleNumberTest {
    private SingleNumber singleNumber;
    @BeforeEach
    void setUp() {
        singleNumber = new SingleNumber();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        int[] nums = {2,2,1};

        int expected = 1;
        int actual = singleNumber.singleNumber(nums);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        int[] nums = {4,1,2,1,2};

        int expected = 4;
        int actual = singleNumber.singleNumber(nums);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        int[] nums = {1};

        int expected = 1;
        int actual = singleNumber.singleNumber(nums);

        assertEquals(expected, actual);
    }
}

