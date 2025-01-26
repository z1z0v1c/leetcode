/**
 * 79 - Medium
 *
 * Given an m x n grid of characters board and a string word, return true if word exists in the grid.
 * The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally
 * or vertically neighboring. The same letter cell may not be used more than once.
 *
 * Example 1:
 *
 * 	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
 * 	Output: true
 *
 * Example 2:
 *
 *  Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
 *  Output: true
 *
 * Example 3:
 *
 * 	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
 * 	Output: false
 *
 * Constraints:
 *
 * 	- m == board.length
 * 	- n = board[i].length
 * 	- 1 <= m, n <= 6
 * 	- 1 <= word.length <= 15
 * 	- board and word consists of only lowercase and uppercase English letters.
 *
 * Follow up: Could you use search pruning to make your solution faster with a larger board?
 */
package wordsearch

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			exist := checkAround(board, word, i, j, 0)
			if exist {
				return true
			}
		}
	}

	return false
}

func checkAround(board [][]byte, word string, i, j, k int) bool {
	if k == len(word) {
		return true
	} else if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[k] {
		return false
	}

	temp := board[i][j]
	board[i][j] = ' '

	if checkAround(board, word, i, j+1, k+1) || checkAround(board, word, i, j-1, k+1) ||
		checkAround(board, word, i+1, j, k+1) || checkAround(board, word, i-1, j, k+1) {
		return true
	}

	board[i][j] = temp

	return false
}
