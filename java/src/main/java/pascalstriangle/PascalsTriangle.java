/**
 * 118 - Easy
 * <p>
 * Given an integer numRows, return the first numRows of Pascal's triangle.
 * In Pascal's triangle, each number is the sum of the two numbers directly above it
 * <p>
 * Example 1:
 *      Input: numRows = 5
 *      Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 * <p>
 * Example 2:
 *      Input: numRows = 1
 *      Output: [[1]]
 * <p>
 * Constraints:
 *      - 1 <= numRows <= 30
 */
package pascalstriangle;

import java.util.ArrayList;
import java.util.List;

public class PascalsTriangle {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> pyramid = new ArrayList<>();
        List<Integer> nullRow = new ArrayList<>();
        nullRow.add(1);

        for (var i = 0; i < numRows; i++) {
            List<Integer> previousRow = i == 0 ? nullRow : pyramid.get(i - 1);
            List<Integer> row = new ArrayList<>();

            for (int j = 0; j <= i; j++) {
                int first = j == 0 ? 0 : previousRow.get(j - 1);
                int second = j == previousRow.size() ? 0 : previousRow.get(j);

                row.add(first + second);
            }

            pyramid.add(row);
        }

        return pyramid;
    }
}

