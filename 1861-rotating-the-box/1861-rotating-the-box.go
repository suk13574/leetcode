func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])

	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)

		for j := 0; j < m; j++ {
			res[i][j] = '.'
		}
	}

	for j := 0; j < m; j++ {
		i := n - 1

		x := m - 1 - j
		y := n - 1

		for y >= 0 && y < n {
			if boxGrid[x][y] == '#' {
				res[i][j] = '#'
				i--
			} else if boxGrid[x][y] == '*' {
				i = y
				res[i][j] = '*'
				i--
			}

			y--
		}
	}

	return res
}