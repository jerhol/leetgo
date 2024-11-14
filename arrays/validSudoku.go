package arrays

// Time: O(n^2)
// Space: O(n^2)

func ValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	sqrs := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		sqrs[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c]
			s := (r/3)*3 + c/3

			if rows[r][val] || cols[c][val] || sqrs[s][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			sqrs[s][val] = true
		}
	}

	return true
}
