func resultArray(nums []int, k int) []int64 {
	n := len(nums)

	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k)
	}

	dp[0][nums[0]%k] = int64(1)

	for i := 1; i < n; i++ {
		v := nums[i]
		dp[i][v%k]++

		for r := 0; r < k; r++ {
			nr := (r * v) % k
			dp[i][nr] += dp[i-1][r]
		}
	}

	res := make([]int64, k)
	for r := 0; r < k; r++ {
		for i := 0; i < n; i++ {
			res[r] += dp[i][r]
		}
	}

	return res
}