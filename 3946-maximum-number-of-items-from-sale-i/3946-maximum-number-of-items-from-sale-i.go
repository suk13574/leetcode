func maximumSaleItems(items [][]int, budget int) int {
	n := len(items)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, budget+1)
	}

	for k := 0; k < n; k++ {
		factor := items[k][0]
		price := items[k][1]
		free := 0

		for i := 0; i < n; i++ {
			if i == k {
				continue
			}

			factor2 := items[i][0]
			if factor2%factor == 0 {
				free++
			}
		}

		for b := 0; b <= budget; b++ {
			dp[k+1][b] = dp[k][b]

			if b >= price {
				dp[k+1][b] = max(
					dp[k+1][b],
					dp[k][b-price]+1+free,
					dp[k+1][b-price]+1,
				)
			}
		}
	}

	return dp[n][budget]
}