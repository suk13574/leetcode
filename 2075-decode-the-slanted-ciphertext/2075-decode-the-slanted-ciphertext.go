func decodeCiphertext(encodedText string, rows int) string {
	if len(encodedText) == 0 {
		return ""
	}

	cols := len(encodedText) / rows

	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]byte, cols)
	}

	for i := 0; i < len(encodedText); i++ {
		r := i / cols
		c := i % cols

		b := encodedText[i]
		grid[r][c] = b
	}

	res := []byte{}
	for startCol := 0; startCol < cols; startCol++ {
		r, c := 0, startCol
		for r < rows && c < cols {
			res = append(res, grid[r][c])
			r++
			c++
		}
	}

	return strings.TrimRight(string(res), " ")
}