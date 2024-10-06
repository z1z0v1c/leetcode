package pascalstriangle;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class PascalsTriangleTest {
    private PascalsTriangle pascalsTriangle;

    @BeforeEach
    void setUp() {
        pascalsTriangle = new PascalsTriangle();
    }

    @Test
    @DisplayName("Five rows")
    void testGenerateFiveRows() {
        int rowIndex = 5;
        var expected = new ArrayList<>(
                Arrays.asList(
                        List.of(1),
                        Arrays.asList(1, 1),
                        Arrays.asList(1, 2, 1),
                        Arrays.asList(1, 3, 3, 1),
                        Arrays.asList(1, 4, 6, 4, 1)
                )
        );
        var actual = pascalsTriangle.generate(rowIndex);

        assertEquals(expected, actual, "Should return [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]");
    }

    @Test
    @DisplayName("One row")
    void testGenerateOneRow() {
        int rowIndex = 1;
        var expected = new ArrayList<>(
                List.of(List.of(1))
        );
        var actual = pascalsTriangle.generate(rowIndex);

        assertEquals(expected, actual, "Should return [[1]]");
    }
}
