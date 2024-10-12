package validpalindrome;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ValidPalindromeTest {
    private ValidPalindrome validPalindrome;

    @BeforeEach
    void setUp() {
        validPalindrome = new ValidPalindrome();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        String s = "A man, a plan, a canal: Panama";

        assertTrue(validPalindrome.isPalindrome(s));
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        String s = "race a car";

        assertFalse(validPalindrome.isPalindrome(s));
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        String s = "  ";

        assertTrue(validPalindrome.isPalindrome(s));
    }

    @Test
    @DisplayName("Example 4.")
    void testExampleFour() {
        String s = "0P";

        assertFalse(validPalindrome.isPalindrome(s));
    }
}
