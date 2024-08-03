import org.example.StringToInteger;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class StringToIntegerTests {
    private StringToInteger stringToInteger;

    @BeforeEach
    void setUp() {
        StringToIntegerTests stringToInteger = new StringToIntegerTests();
    }

    @Test
    @DisplayName("Simple multiplication should work")
    void testMyAtoiEmptyString() {
        assertEquals(0, stringToInteger.myAtoi(""),
                "Regular multiplication should work");
    }

//    @RepeatedTest(5)
//    @DisplayName("Ensure correct handling of zero")
//    void testMultiplyWithZero() {
//        assertEquals(0, calculator.multiply(0, 5), "Multiple with zero should be zero");
//        assertEquals(0, calculator.multiply(5, 0), "Multiple with zero should be zero");
//    }
}