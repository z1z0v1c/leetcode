package smallestpositiveinteger;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SmallestPositiveIntegerTest {
    private SmallestPositiveInteger smallestPositiveInteger;

    @BeforeEach
    void setUp() {
        smallestPositiveInteger = new SmallestPositiveInteger();
    }

    @Test
    @DisplayName("Should return 5")
    void testSmallestPositiveInteger5() {
        int[] A = new int[]{1, 3, 6, 4, 1, 2};
        assertEquals(5, smallestPositiveInteger.smallestPositiveInteger(A),
                "Should return 5");
    }

    @Test
    @DisplayName("Should return 4")
    void testSmallestPositiveInteger4() {
        int[] A = new int[]{1, 2, 3};
        assertEquals(4, smallestPositiveInteger.smallestPositiveInteger(A),
                "Should return 4");
    }

    @Test
    @DisplayName("All negative")
    void testSmallestPositiveIntegerAllNegative() {
        int i = -1;
        int j = -3;
        int k = -7;
        int[] A = new int[]{i, j, k};

        assertEquals(1, smallestPositiveInteger.smallestPositiveInteger(A),
                "Should return 1");
    }
}
