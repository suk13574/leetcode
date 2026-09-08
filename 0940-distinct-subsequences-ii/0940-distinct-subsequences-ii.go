func distinctSubseqII(s string) int {
	const MOD int64 = 1_000_000_007

	prev := make([]int, 26)
	for i := 0; i < 26; i++ {
		prev[i] = -1
	}

	dp := make([]int64, len(s)+1)
	dp[0] = int64(1)

	for i, ch := range s {
		pos := ch - 'a'

		dp[i+1] = (dp[i] * 2) % MOD

		if prev[pos] != -1 {
			dp[i+1] = (dp[i+1] - dp[prev[pos]] + MOD) % MOD
		}

		prev[pos] = i
	}

	return int((dp[len(s)] - 1 + MOD) % MOD)
}