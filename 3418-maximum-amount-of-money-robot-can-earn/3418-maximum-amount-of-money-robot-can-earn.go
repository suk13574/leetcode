func maximumAmount(coins [][]int) int {
	const NEG = -1 << 60
	r, c := len(coins), len(coins[0])

	dp := make([][][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([][]int, c)
		for j := 0; j < c; j++ {
			dp[i][j] = []int{NEG, NEG, NEG}
		}
	}

	if coins[0][0] >= 0 {
		dp[0][0][0] = coins[0][0]
	} else {
		dp[0][0][0] = coins[0][0]
		dp[0][0][1] = 0
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				continue
			}

			for k := 0; k < 3; k++ {
				if i > 0 && dp[i-1][j][k] != NEG {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k]+coins[i][j])
				}
				if j > 0 && dp[i][j-1][k] != NEG {
					dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k]+coins[i][j])
				}

				if coins[i][j] < 0 && k > 0 {
					if i > 0 && dp[i-1][j][k-1] != NEG {
						dp[i][j][k] = max(dp[i][j][k], dp[i-1][j][k-1])
					}
					if j > 0 && dp[i][j-1][k-1] != NEG {
						dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k-1])
					}
				}
			}
		}
	}

	return max(dp[r-1][c-1][0], max(dp[r-1][c-1][1], dp[r-1][c-1][2]))
}