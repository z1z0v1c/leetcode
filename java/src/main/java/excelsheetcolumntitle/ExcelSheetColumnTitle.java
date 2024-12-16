/**
 * 168 - Easy
 * <p>
 * Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
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
 *      Input: columnNumber = 1
 *      Output: "A"
 * <p>
 * Example 2:
 *      Input: columnNumber = 28
 *      Output: "AB"
 * <p>
 * Example 3:
 *      Input: columnNumber = 701
 *      Output: "ZY"
 * <p>
 * Constraints:
 *      - 1 <= columnNumber <= 231 - 1
 */
package excelsheetcolumntitle;

public class ExcelSheetColumnTitle {
    public String convertToTitle(int columnNumber) {
        StringBuilder columnTitle = new StringBuilder();

        while (columnNumber > 0) {
            int offset = columnNumber % 26 - 1;

            if (offset == - 1) {
                columnTitle.append('Z');
                columnNumber = columnNumber / 26 - 1;
            } else {
                columnTitle.append(Character.toString('A' + offset));
                columnNumber /= 26;
            }
        }

        return  columnTitle.reverse().toString();
    }
}

