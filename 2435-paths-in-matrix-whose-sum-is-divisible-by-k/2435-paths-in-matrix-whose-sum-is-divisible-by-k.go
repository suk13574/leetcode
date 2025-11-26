func numberOfPaths(grid [][]int, k int) int {
	var MOD = 1000000007

	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k)
		}
	}

	dp[0][0][grid[0][0]%k]++

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			for r := 0; r < k; r++ {
				if i > 0 {
					newR := (r + val) % k
					dp[i][j][newR] += dp[i-1][j][r]
					dp[i][j][newR] %= MOD

				}

				if j > 0 {
					newR := (r + val) % k
					dp[i][j][newR] += dp[i][j-1][r]
					dp[i][j][newR] %= MOD
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}