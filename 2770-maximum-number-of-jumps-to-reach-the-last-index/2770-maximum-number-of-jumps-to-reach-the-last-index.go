func maximumJumps(nums []int, target int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
            if dp[i] == -1 {
                continue
            }
			gap := nums[i] - nums[j]
			if abs(gap) <= target {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}