func myPow(n, x int) int {
	res := 1
	base := n

	for x > 0 {
		if x&1 == 1 {
			res = res * base
		}
		base = base * base
		x = x >> 1
	}
	return res
}

func numberOfWays(n int, x int) int {
	const MOD = 1000000007

	powers := []int{}
	for N := 1; ; N++ {
		p := myPow(N, x)
		if p > n {
			break
		}
		powers = append(powers, p)
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for _, power := range powers {
		for c := n; c >= power; c-- {
			dp[c] = dp[c] + dp[c-power]
		}
	}

	return dp[n] % MOD
}