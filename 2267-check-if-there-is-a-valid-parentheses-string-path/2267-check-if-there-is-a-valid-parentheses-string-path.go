func hasValidPath(grid [][]byte) bool {
	if grid[0][0] == ')' {
		return false
	}

	m, n := len(grid), len(grid[0])

	dp := make([][][]bool, m)

	for i := 0; i < m; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, m+n)
		}
	}

	dp[0][0][1] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for bal := 0; bal < m+n; bal++ {
				if !dp[i][j][bal] {
					continue
				}

				if i+1 < m {
					nextBal := bal

					if grid[i+1][j] == '(' {
						nextBal++
					} else {
						nextBal--
					}

					if nextBal >= 0 {
						dp[i+1][j][nextBal] = true
					}
				}

				if j+1 < n {
					nextBal := bal

					if grid[i][j+1] == '(' {
						nextBal++
					} else {
						nextBal--
					}

					if nextBal >= 0 {
						dp[i][j+1][nextBal] = true
					}
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}