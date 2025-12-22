func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	dp := make([]int, m)
	best := 1

	for j := 0; j < m; j++ {
		dp[j] = 1
		for i := 0; i < j; i++ {
			ok := true
			for r := 0; r < n; r++ {
				if strs[r][i] > strs[r][j] {
					ok = false
					break
				}
			}
			if ok && dp[i]+1 > dp[j] {
				dp[j] = dp[i] + 1
			}
		}
		best = max(best, dp[j])
	}

	return m - best
}
