package vaildsudoku

func isValidSudoku(board [][]byte) bool {
	cols := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	squares := make([][]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		cols[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
	}

	for i := 0; i < 3; i++ {
		squares[i] = make([]map[byte]bool, 3)
		for j := 0; j < 3; j++ {
			squares[i][j] = make(map[byte]bool)
		}
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			num := board[r][c]
			if rows[r][num] || cols[c][num] || squares[r/3][c/3][num] {
				return false
			}
			rows[r][num] = true
			cols[c][num] = true
			squares[r/3][c/3][num] = true
		}
	}
	return true
}
