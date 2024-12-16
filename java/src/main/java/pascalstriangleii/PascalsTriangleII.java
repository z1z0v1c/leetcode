/**
 * 119 - Easy
 * <p>
 * Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
 * In Pascal's triangle, each number is the sum of the two numbers directly above it
 * <p>
 * Example 1:
 *      Input: rowIndex = 3
 *      Output: [1,3,3,1]
 * <p>
 * Example 2:
 *      Input: rowIndex = 0
 *      Output: [1]
 * <p>
 * Example 3:
 *      Input: rowIndex = 1
 *      Output: [1,1]
 * <p>
 * Constraints:
 *      - 0 <= rowIndex <= 33
 */
package pascalstriangleii;

import java.util.ArrayList;
import java.util.List;

public class PascalsTriangleII {
    public List<Integer> getRow(int rowIndex) {
        List<Integer> row = new ArrayList<>();
        List<Integer> previousRow = new ArrayList<>();
        previousRow.add(1);

        for (int i = 0; i < rowIndex; i++) {
            row.clear();

            for (int j = 0; j <= previousRow.size(); j++) {
                int previous = j == 0 ? 0 : previousRow.get(j - 1);
                int next = j == previousRow.size() ? 0 : previousRow.get(j);
                row.add(previous + next);
            }

            previousRow.clear();
            previousRow.addAll(row);
        }

        return row.isEmpty() ? previousRow : row;
    }
}

