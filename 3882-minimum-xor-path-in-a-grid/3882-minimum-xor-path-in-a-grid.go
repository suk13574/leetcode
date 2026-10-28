func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][][]bool, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, 1024)
		}
	}

	dp[0][0][grid[0][0]] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			v := grid[i][j]

			for x := 0; x < 1024; x++ {
				if i > 0 && dp[i-1][j][x] {
					dp[i][j][x^v] = true
				}

				if j > 0 && dp[i][j-1][x] {
					dp[i][j][x^v] = true
				}
			}
		}
	}

	for x := 0; x < 1024; x++ {
		if dp[m-1][n-1][x] {
			return x
		}
	}

	return -1
}