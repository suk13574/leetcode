func uniquePaths(grid [][]int) int {
    const (
        Down  = 0
        Right = 1
        MOD   = 1_000_000_007
    )

	m, n := len(grid), len(grid[0])

	dp := make([][][2]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][2]int, n)
	}

	dp[0][0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			down := dp[i][j][Down]
			right := dp[i][j][Right]

			if grid[i][j] == 0 {
                total := (down + right) % MOD
                down, right = total, total
            } else {
				down, right = right, down
            }

			if i+1 < m {
				dp[i+1][j][0] = (dp[i+1][j][0] + down) % MOD
			}

			if j+1 < n {
				dp[i][j+1][1] = (dp[i][j+1][1] + right) % MOD
			}
		}
	}

	return (dp[m-1][n-1][0] + dp[m-1][n-1][1]) % MOD
}