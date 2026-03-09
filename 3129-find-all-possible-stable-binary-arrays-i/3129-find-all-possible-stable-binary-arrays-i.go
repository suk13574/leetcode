func numberOfStableArrays(zero int, one int, limit int) int {
	const MOD = 1_000_000_007

	// dp[z][o][last][run]:
	// z    = number of 0s used so far
	// o    = number of 1s used so far
	// last = last placed value (0 or 1)
	// run  = current consecutive length of `last`
	dp := make([][][][]int, zero+1)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][][]int, one+1)
		for j := 0; j <= one; j++ {
			dp[i][j] = make([][]int, 2)
			for last := 0; last < 2; last++ {
				dp[i][j][last] = make([]int, limit+1)
			}
		}
	}

	if zero > 0 {
		dp[1][0][0][1] = 1
	}
	if one > 0 {
		dp[0][1][1][1] = 1
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			for last := 0; last < 2; last++ {
				for cnt := 1; cnt <= limit; cnt++ {
					cur := dp[i][j][last][cnt]

					if cur == 0 {
						continue
					}

					if last == 0 {
						if i < zero && cnt < limit {
							dp[i+1][j][0][cnt+1] = (dp[i+1][j][0][cnt+1] + cur) % MOD
						}

						if j < one {
							dp[i][j+1][1][1] = (dp[i][j+1][1][1] + cur) % MOD
						}
					} else {
						if j < one && cnt < limit {
							dp[i][j+1][1][cnt+1] = (dp[i][j+1][1][cnt+1] + cur) % MOD
						}

						if i < zero {
							dp[i+1][j][0][1] = (dp[i+1][j][0][1] + cur) % MOD
						}
					}
				}
			}
		}
	}

	res := 0
	for cnt := 1; cnt <= limit; cnt++ {
		res = (res + dp[zero][one][0][cnt]) % MOD
		res = (res + dp[zero][one][1][cnt]) % MOD
	}
	return res
}