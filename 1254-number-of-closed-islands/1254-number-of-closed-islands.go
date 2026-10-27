func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	moves := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}

		if grid[x][y] != 0 {
			return
		}

		grid[x][y] = 1

		for _, move := range moves {
			dfs(x+move[0], y+move[1])
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	res := 0

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}