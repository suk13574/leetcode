func numberOfSets(n int, k int) int {
    const MOD = 1_000_000_007

    // dp[k][n]
    dp := make([][]int, k+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, n+1)
    }

    for i := 0; i < n; i++ {
        dp[0][i] = 1
    }    

   // prefix[j][i] = sum of dp[j][0...i]
    prefix := make([][]int, k+2)
    for i := 0; i < len(prefix); i++ {
        prefix[i] = make([]int, n+2)
    }

    prefix[0][0] = dp[0][0]
    for i := 1; i < n; i++ {
        prefix[0][i] = (prefix[0][i-1] + dp[0][i]) % MOD
    }

    for j := 1; j <= k; j++ {
        prefix[j][0] = dp[j][0]

        for i := 1; i < n; i++ {
            dp[j][i] = (dp[j][i-1] + prefix[j-1][i-1]) % MOD
            prefix[j][i] = (prefix[j][i-1] + dp[j][i]) % MOD
        }
    }

    return dp[k][n-1]
}