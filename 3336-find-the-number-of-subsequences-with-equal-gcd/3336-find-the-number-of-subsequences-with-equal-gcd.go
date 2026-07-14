func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount(nums []int) int {
	const MOD int64 = 1_000_000_007

	maxNum := 0
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	dp := make([][]int64, maxNum+1)
	for i := range dp {
		dp[i] = make([]int64, maxNum+1)
	}

	dp[0][0] = 1

	for _, num := range nums {
		next := make([][]int64, maxNum+1)
		for i := range dp {
			next[i] = make([]int64, maxNum+1)
		}

		for g1 := 0; g1 <= maxNum; g1++ {
			for g2 := 0; g2 <= maxNum; g2++ {
				count := dp[g1][g2]
				if count == 0 {
					continue
				}

				next[g1][g2] = (next[g1][g2] + count) % MOD

				nextG1 := gcd(num, g1)
				next[nextG1][g2] = (next[nextG1][g2] + count) % MOD

				nextG2 := gcd(num, g2)
				next[g1][nextG2] = (next[g1][nextG2] + count) % MOD
			}
		}

		dp = next
	}

	var res int64
	for g := 1; g <= maxNum; g++ {
		res = (res + dp[g][g]) % MOD
	}

	return int(res)
}