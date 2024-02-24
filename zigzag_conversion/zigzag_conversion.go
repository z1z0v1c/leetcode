// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
// (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:

// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000

package zigzagconversion

func convert(s string, numRows int) string {
	if numRows >= len(s) {
		return s
	}

	str := ""

	char := 0
	slice := make([][]byte, 0)

	i := 0
while:
	for {
		row := make([]byte, numRows)
		slice = append(slice, row)
		for j := 0; j < numRows; j++ {
			slice[i][j] = s[char]
			char++
			if char == len(s) {
				break while
			}
		}
		i++
	}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[j][i] != 0 {
				str += string(slice[j][i])
			}
		}
	}
	return str
}
