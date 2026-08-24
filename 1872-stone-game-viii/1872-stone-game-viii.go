func stoneGameVIII(stones []int) int {
    n := len(stones)

    prefix := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + stones[i]
    }
    
    dp := prefix[n]
    
    for i := n-1; i > 1; i-- {
        dp = max(dp, prefix[i]-dp)
    }

    return dp
}