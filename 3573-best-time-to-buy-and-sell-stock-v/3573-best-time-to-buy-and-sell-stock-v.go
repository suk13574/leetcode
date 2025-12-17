func maximumProfit(prices []int, k int) int64 {
	const NEG int64 = -(1 << 60)

	n := len(prices)

	dp := make([][][]int64, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int64, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = []int64{NEG, NEG, NEG} // state: do nothing, long, short
		}
	}

	for kk := 0; kk <= k; kk++ {
		dp[n][kk][0] = 0 // last day, do nothing
	}

	for i := n - 1; i >= 0; i-- {
		p := int64(prices[i])
		for kk := 0; kk <= k; kk++ {
			best0 := dp[i+1][kk][0]              // do nothing
			best0 = max(best0, dp[i+1][kk][1]-p) // start long (buy)
			best0 = max(best0, dp[i+1][kk][2]+p) // start short (sell)
			dp[i][kk][0] = best0

			// state 1: long holding
			best1 := dp[i+1][kk][1] // carry
			if kk > 0 {
				best1 = max(best1, dp[i+1][kk-1][0]+p) // close long (sell)
			}
			dp[i][kk][1] = best1

			// state 2: short holding
			best2 := dp[i+1][kk][2] // carry
			if kk > 0 {
				best2 = max(best2, dp[i+1][kk-1][0]-p) // close short (buyback)
			}
			dp[i][kk][2] = best2
		}
	}

	return dp[0][k][0]
}