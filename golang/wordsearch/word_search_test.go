package wordsearch

import "testing"

func TestExist(t *testing.T) {
	// Example one
	word := "ABCCED"
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	if !exist(board, word) {
		t.Errorf("exist(%#v, %s) should return true for example 1.", word, board)
	}

	// Example two
	word = "SEE"
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	if !exist(board, word) {
		t.Errorf("exist(%#v, %s) should return true for example 2.", word, board)
	}

	// Example three
	word = "ABCB"
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	if exist(board, word) {
		t.Errorf("exist(%#v, %s) should return false for example 3.", word, board)
	}
}
