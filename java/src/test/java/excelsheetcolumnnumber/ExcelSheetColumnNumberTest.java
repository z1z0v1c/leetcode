package excelsheetcolumnnumber;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ExcelSheetColumnNumberTest {
    private ExcelSheetColumnNumber excelSheetColumnNumber;

    @BeforeEach
    void setUp() {
        excelSheetColumnNumber = new ExcelSheetColumnNumber();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        String columnTitle = "A";

        int expected = 1;
        int actual = excelSheetColumnNumber.titleToNumber(columnTitle);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        String columnTitle = "AB";

        int expected = 28;
        int actual = excelSheetColumnNumber.titleToNumber(columnTitle);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        String columnTitle = "ZY";

        int expected = 701;
        int actual = excelSheetColumnNumber.titleToNumber(columnTitle);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 4.")
    void testExampleFour() {
        String columnTitle = "FXSHRXW";

        int expected = 2147483647;
        int actual = excelSheetColumnNumber.titleToNumber(columnTitle);

        assertEquals(expected, actual);
    }
}
