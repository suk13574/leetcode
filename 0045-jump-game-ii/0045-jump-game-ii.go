func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)

    for i := 1; i < n; i++ {
        dp[i] = n+1
    }

    for i := 0; i < n; i++ {
        for j := 1; j <= nums[i] && i+j < n; j++ {
            dp[i+j] = min(dp[i+j], dp[i]+1)
        }
    }

    return dp[n-1]
}