/**
 * 36 - Medium
 *
 * Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
 *
 * 	- Each row must contain the digits 1-9 without repetition.
 * 	- Each column must contain the digits 1-9 without repetition.
 * 	- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
 *
 * Note:
 *
 * 	A Sudoku board (partially filled) could be valid but is not necessarily solvable.
 * 	Only the filled cells need to be validated according to the mentioned rules.
 *
 * Example 1:
 *
 * 	Input: board =
 *
 * 		[["5","3",".",".","7",".",".",".","."]
 * 		,["6",".",".","1","9","5",".",".","."]
 * 		,[".","9","8",".",".",".",".","6","."]
 * 		,["8",".",".",".","6",".",".",".","3"]
 * 		,["4",".",".","8",".","3",".",".","1"]
 * 		,["7",".",".",".","2",".",".",".","6"]
 * 		,[".","6",".",".",".",".","2","8","."]
 * 		,[".",".",".","4","1","9",".",".","5"]
 * 		,[".",".",".",".","8",".",".","7","9"]]
 *
 * 	Output: true
 *
 * Example 2:
 *
 * 	Input: board =
 *
 * 		[["8","3",".",".","7",".",".",".","."]
 * 		,["6",".",".","1","9","5",".",".","."]
 * 		,[".","9","8",".",".",".",".","6","."]
 * 		,["8",".",".",".","6",".",".",".","3"]
 * 		,["4",".",".","8",".","3",".",".","1"]
 * 		,["7",".",".",".","2",".",".",".","6"]
 * 		,[".","6",".",".",".",".","2","8","."]
 * 		,[".",".",".","4","1","9",".",".","5"]
 * 		,[".",".",".",".","8",".",".","7","9"]]
 *
 * 	Output: false
 *
 * 	Explanation:
 *
 * 		Same as Example 1, except with the 5 in the top left corner being modified to 8.
 * 		Since there are two 8's in the top left 3x3 sub-box, it is invalid.
 *
 * Constraints:
 *
 * 	- board.length == 9
 * 	- board[i].length == 9
 * 	- board[i][j] is a digit 1-9 or '.'.
 */
package validsudoku

func isValidSudoku(board [][]byte) bool {
	// Check rows and columns
	for i := 0; i < len(board); i++ {
		row := make(map[byte]int, len(board))
		col := make(map[byte]int, len(board))

		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				if board[i][j] < '1' || board[i][j] > '9' {
					return false
				}

				row[board[i][j]]++
				if row[board[i][j]] > 1 {
					return false
				}
			}

			if board[j][i] != '.' {
				if board[j][i] < '1' || board[j][i] > '9' {
					return false
				}

				col[board[j][i]]++
				if col[board[j][i]] > 1 {
					return false
				}
			}
		}
	}

	// Check all squares
	for k := 0; k < 3; k++ {
		for z := 0; z < 3; z++ {
			square := make(map[byte]int, len(board))

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[k*3+i][z*3+j] == '.' {
						continue
					}

					if board[k*3+i][z*3+j] < '1' || board[k*3+i][z*3+j] > '9' {
						return false
					}

					square[board[k*3+i][z*3+j]]++
					if square[board[k*3+i][z*3+j]] > 1 {
						return false
					}
				}
			}
		}
	}

	return true
}
