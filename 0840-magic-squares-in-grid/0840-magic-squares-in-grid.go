func numMagicSquaresInside(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	if n < 2 || m < 2 {
		return 0
	}

	res := 0
	for i := 0; i+2 < n; i++ {
	loop:
		for j := 0; j+2 < m; j++ {
			hash := make([]bool, 10)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					v := grid[i+x][j+y]
					if v <= 0 || v > 9 || hash[v] {
						continue loop
					}
					hash[v] = true
				}
			}

			d1 := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
			d2 := grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]
			if d1 != d2 {
				continue loop
			}

			r1 := grid[i][j] + grid[i][j+1] + grid[i][j+2]
			r2 := grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2]
			r3 := grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if r1 != r2 || r1 != r3 {
				continue loop
			}

			// check sum of column
			c1 := grid[i][j] + grid[i+1][j] + grid[i+2][j]
			c2 := grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1]
			c3 := grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]
			if c1 != c2 || c1 != c3 {
				continue loop
			}

			res++
		}
	}

	return res
}