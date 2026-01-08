func maxDotProduct(nums1 []int, nums2 []int) int {
    const NEG = -1 << 30

    n, m := len(nums1), len(nums2)

    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j <= m; j++ {
            dp[i][j] = NEG
        }
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            prod := nums1[i-1] * nums2[j-1]
            
            // 1) Start a new subsequence 
            best := prod

            // 2) Extend
            if dp[i-1][j-1]+prod > best {
                best = dp[i-1][j-1]+prod
            }

            // Skip nums1[i-1]
            if dp[i-1][j] > best {
                best = dp[i-1][j]
            }

            // Skip nums2[j-1]
            if dp[i][j-1] > best {
                best = dp[i][j-1]
            }
            
            dp[i][j] = best
        }
    }

    return dp[n][m]
}