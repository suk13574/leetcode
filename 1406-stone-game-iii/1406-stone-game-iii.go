func stoneGameIII(stoneValue []int) string {
    n := len(stoneValue)
    
    dp := make([]int, n+3)

    for i := n-1; i >= 0; i-- {
        v := stoneValue[i]

        // take 1 stone
        maxValue := v - dp[i+1]

        // take 2 stone 
        if i+1 < n {
            maxValue = max(maxValue, v + stoneValue[i+1] - dp[i+2])
        }

        // take 3 stone
        if i+2 < n {
            maxValue = max(maxValue, v + stoneValue[i+1] + stoneValue[i+2] - dp[i+3])
        }

        dp[i] = maxValue
    }

    if dp[0] > 0 {
        return "Alice"
    } else if dp[0] < 0 {
        return "Bob"
    } else {
        return "Tie"
    }
}