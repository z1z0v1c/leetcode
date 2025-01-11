/**
 * 6 - Medium
 * 
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
 * (you may want to display this pattern in a fixed font for better legibility)
 * 
 * 	P   A   H   N
 * 	A P L S I I G
 * 	Y   I   R
 * 
 * And then read line by line: "PAHNAPLSIIGYIR"
 * 
 * Write the code that will take a string and make this conversion given a number of rows:
 * 
 * 	string convert(string s, int numRows);
 * 
 * Example 1:
 * 
 * 	Input: s = "PAYPALISHIRING", numRows = 3
 * 	Output: "PAHNAPLSIIGYIR"
 * 
 * Example 2:
 * 
 * 	Input: s = "PAYPALISHIRING", numRows = 4
 * 	Output: "PINALSIGYAHRPI"
 * 
 * 	Explanation:
 * 
 * 		P     I    N
 * 		A   L S  I G
 * 		Y A   H R
 * 		P     I
 * 
 * Example 3:
 * 
 * 	Input: s = "A", numRows = 1
 * 	Output: "A"
 * 
 * Constraints:
 * 
 * 	- 1 <= s.length <= 1000
 * 	- s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 	- 1 <= numRows <= 1000
 */
package zigzagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	converted := ""
	index := 0

	// TODO - use a slice instead of matrix
	matrix := make([][]byte, 1)
	matrix[0] = make([]byte, numRows)

	i := 0
	j := 0

	endOfRow := false

	for index < len(s) {
		matrix[i][j] = s[index]
		index++

		if j == numRows-1 {
			endOfRow = true
		}

		if endOfRow {
			row := make([]byte, numRows)
			matrix = append(matrix, row)

			j--
			i++
		} else {
			j++
		}

		if j == 0 {
			endOfRow = false
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] != 0 {
				converted += string(matrix[j][i])
			}
		}
	}

	return converted
}

