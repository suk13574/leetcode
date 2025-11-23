func maxSumDivThree(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{-1, -1, -1}
	}

	dp[0][0] = 0
	dp[0][nums[0]%3] = nums[0]

	for i := 1; i < n; i++ {
		for r := 0; r < 3; r++ {
			dp[i][r] = dp[i-1][r]
		}

		for r := 0; r < 3; r++ {
			if dp[i-1][r] == -1 {
				continue
			}
			newR := (r + nums[i]) % 3
			sum := dp[i-1][r] + nums[i]
			dp[i][newR] = max(dp[i][newR], sum)
		}
	}

	if dp[n-1][0] == -1 {
		return 0
	}
	return dp[n-1][0]
}