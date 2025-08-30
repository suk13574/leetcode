func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]uint16

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}

			bit := uint16(1) << (ch - '1')
			boxIdx := (i/3)*3 + (j / 3)

			if (rows[i]&bit) != 0 || (cols[j]&bit) != 0 || (boxes[boxIdx]&bit) != 0 {
				return false
			}

			rows[i] |= bit
			cols[j] |= bit
			boxes[boxIdx] |= bit
		}
	}
	return true
}