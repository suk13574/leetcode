func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			minScore := math.MaxInt64
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[k]*values[j]
				minScore = min(minScore, score)
			}
			dp[i][j] = minScore
		}
	}

	return dp[0][n-1]
}