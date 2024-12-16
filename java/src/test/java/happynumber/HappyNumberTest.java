package happynumber;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class HappyNumberTest {
    private HappyNumber happyNumber;

    @BeforeEach
    void setUp() {
        happyNumber = new HappyNumber();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        int n = 19;

        assertTrue(happyNumber.isHappy(n));
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        int n = 2;

        assertFalse(happyNumber.isHappy(n));
    }
}

