package majorityelement;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MajorityElementTest {
    private MajorityElement majorityElement;

    @BeforeEach
    void setUp() {
        majorityElement = new MajorityElement();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        int[] nums = {3, 2, 3};

        int expected = 3;
        int actual = majorityElement.majorityElement(nums);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        int[] nums = {2, 2, 1, 1, 1, 2, 2};

        int expected = 2;
        int actual = majorityElement.majorityElement(nums);

        assertEquals(expected, actual);
    }
}
