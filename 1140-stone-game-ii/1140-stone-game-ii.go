func stoneGameII(piles []int) int {
	n := len(piles)

	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for M := 1; M <= n; M++ {
			if 2*M >= n-i {
				dp[i][M] = suffix[i]
				continue
			}

			for X := 1; X <= 2*M; X++ {
				nextM := max(M, X)
				current := suffix[i] - dp[i+X][nextM]

				dp[i][M] = max(dp[i][M], current)
			}
		}
	}

	return dp[0][1]
}