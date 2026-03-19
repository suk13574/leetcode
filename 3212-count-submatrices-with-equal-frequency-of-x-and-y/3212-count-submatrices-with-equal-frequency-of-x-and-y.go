func numberOfSubmatrices(grid [][]byte) int {
	r, c := len(grid), len(grid[0])

	xPrefix := make([][]int, r+1)
	yPrefix := make([][]int, r+1)

	for i := 0; i < r+1; i++ {
		xPrefix[i] = make([]int, c+1)
		yPrefix[i] = make([]int, c+1)
	}

	res := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v := grid[i][j]
			xPrefix[i+1][j+1] = xPrefix[i+1][j] + xPrefix[i][j+1] - xPrefix[i][j]
			yPrefix[i+1][j+1] = yPrefix[i+1][j] + yPrefix[i][j+1] - yPrefix[i][j]

			if v == 'X' {
				xPrefix[i+1][j+1]++
			} else if v == 'Y' {
				yPrefix[i+1][j+1]++
			}

            if xPrefix[i+1][j+1] > 0 && xPrefix[i+1][j+1] == yPrefix[i+1][j+1] {
                res++
            }
		}
	}

	return res
}