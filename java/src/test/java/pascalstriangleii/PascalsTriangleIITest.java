package pascalstriangleii;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class PascalsTriangleIITest {
    private PascalsTriangleII pascalsTriangleII;

    @BeforeEach
    void setUp() {
        pascalsTriangleII = new PascalsTriangleII();
    }

    @Test
    @DisplayName("Fourth row")
    void testGetRowForFourthRow() {
        int rowIndex = 3;
        var expected = new ArrayList<>(Arrays.asList(1, 3, 3, 1));
        var actual = pascalsTriangleII.getRow(rowIndex);

        assertEquals(expected, actual, "Should return [1, 3, 3, 1]");
    }

    @Test
    @DisplayName("First row")
    void testGetRowForFirstRow() {
        int rowIndex = 0;
        var expected = new ArrayList<>(List.of(1));
        var actual = pascalsTriangleII.getRow(rowIndex);

        assertEquals(expected, actual, "Should return [1]");
    }

    @Test
    @DisplayName("Second row")
    void testGetRowForSecondRow() {
        int rowIndex = 1;
        var expected = new ArrayList<>(Arrays.asList(1, 1));
        var actual = pascalsTriangleII.getRow(rowIndex);

        assertEquals(expected, actual, "Should return [1, 1]");
    }
}
