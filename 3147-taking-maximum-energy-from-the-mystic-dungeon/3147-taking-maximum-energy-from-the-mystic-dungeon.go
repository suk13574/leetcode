func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i-k < 0 {
			dp[i] = energy[i]
		} else {
			dp[i] = max(energy[i], energy[i]+dp[i-k])
		}
	}

	res := dp[n-1]
	for i := n - k - 1; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}
