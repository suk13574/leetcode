func numberOfStableArrays(zero int, one int, limit int) int {
	const MOD = 1_000_000_007

	// dp[x][y][k]
	// x = number of 0s used so far
	// y = number of 1s used so far
	// k = last placed value (0 or 1)
	dp := make([][][]int, zero+1)
	pz := make([][]int, zero+1) // prefix of dp[][?][1] along zero-axis
	po := make([][]int, zero+1) // prefix of dp[?][][0] along one-axis
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, one+1)
		pz[i] = make([]int, one+1)
		po[i] = make([]int, one+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	// initialization dp
	for i := 1; i <= min(limit, zero); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		dp[0][j][1] = 1
	}

	// initialization prefix
	for j := 1; j <= one; j++ {
		pz[0][j] = dp[0][j][1]
	}
	for i := 1; i <= zero; i++ {
		po[i][0] = dp[i][0][0]
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			k := min(i, limit)
			left := i - k
			if left == 0 {
				dp[i][j][0] = pz[i-1][j]
			} else {
				dp[i][j][0] = (pz[i-1][j] - pz[left-1][j] + MOD) % MOD
			}

			k = min(j, limit)
			left = j - k
			if left == 0 {
				dp[i][j][1] = po[i][j-1]
			} else {
				dp[i][j][1] = (po[i][j-1] - po[i][left-1] + MOD) % MOD
			}

			pz[i][j] = (pz[i-1][j] + dp[i][j][1]) % MOD
			po[i][j] = (po[i][j-1] + dp[i][j][0]) % MOD
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % MOD
}