package excelsheetcolumntitle;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ExcelSheetColumnTitleTest {
    private ExcelSheetColumnTitle excelSheetColumnTitle;

    @BeforeEach
    void setUp() {
        excelSheetColumnTitle = new ExcelSheetColumnTitle();
    }

    @Test
    @DisplayName("Example 1.")
    void testExampleOne() {
        var columnNumber = 1;

        var expected = "A";
        var actual = excelSheetColumnTitle.convertToTitle(columnNumber);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 2.")
    void testExampleTwo() {
        var columnNumber = 28;

        var expected = "AB";
        var actual = excelSheetColumnTitle.convertToTitle(columnNumber);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 3.")
    void testExampleThree() {
        var columnNumber = 701;

        var expected = "ZY";
        var actual = excelSheetColumnTitle.convertToTitle(columnNumber);

        assertEquals(expected, actual);
    }

    @Test
    @DisplayName("Example 4.")
    void testExampleFour() {
        var columnNumber = 2147483647;

        var expected = "FXSHRXW";
        var actual = excelSheetColumnTitle.convertToTitle(columnNumber);

        assertEquals(expected, actual);
    }
}

