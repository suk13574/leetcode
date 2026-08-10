func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false

	for i := 1; i <= n; i++ {
		win := false
		for j := 1; j*j <= i; j++ {
			pos := j * j

			if dp[i-pos] == false {
				win = true
				break
			}
		}

		dp[i] = win
	}

	return dp[n]
}