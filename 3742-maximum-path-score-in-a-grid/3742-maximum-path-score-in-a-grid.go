func maxPathScore(grid [][]int, k int) int {
	const NEG = math.MinInt

	m, n := len(grid), len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
			for c := 0; c <= k; c++ {
				dp[i][j][c] = NEG
			}
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			score := grid[i][j]
			cost := 0
			if grid[i][j] != 0 {
				cost = 1
			}

			for c := cost; c <= k; c++ {
				best := NEG

				if i > 0 {
					best = max(best, dp[i-1][j][c-cost])
				}

				if j > 0 {
					best = max(best, dp[i][j-1][c-cost])
				}

				if best != NEG {
					dp[i][j][c] = best + score
				}
			}
		}
	}

	res := NEG
	for c := 0; c <= k; c++ {
		res = max(res, dp[m-1][n-1][c])
	}

	if res == NEG {
		return -1
	}
	return res
}