/**
 * 171 - Easy
 * <p>
 * Given a string columnTitle that represents the column title as appears in an Excel sheet,
 * return its corresponding column number.
 * <p>
 * For example:
 * <p>
 *      A -> 1
 *      B -> 2
 *      C -> 3
 *      ...
 *      Z -> 26
 *      AA -> 27
 *      AB -> 28
 *      ...
 * <p>
 * Example 1:
 *      Input: columnTitle = "A"
 *      Output: 1
 * <p>
 * Example 2:
 *      Input: columnTitle = "AB"
 *      Output: 28
 * <p>
 * Example 3:
 *      Input: columnTitle = "ZY"
 *      Output: 701
 * <p>
 * Constraints:
 *      - 1 <= columnTitle.length <= 7
 *      - columnTitle consists only of uppercase English letters.
 *      - columnTitle is in the range ["A", "FXSHRXW"].
 */
package excelsheetcolumnnumber;

public class ExcelSheetColumnNumber {
    public int titleToNumber(String columnTitle) {
        int columnNumber = 0;
        char[] chars = columnTitle.toCharArray();

        for (int i = chars.length - 1, j = 0; i >= 0; i--, j++) {
            columnNumber += (chars[i] - 'A' + 1) * (int)Math.pow(26, j);
        }

        return columnNumber;
    }
}

