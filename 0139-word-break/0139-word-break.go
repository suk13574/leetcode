func wordBreak(s string, wordDict []string) bool {
    n := len(s)

    word := make(map[string]bool)

    for _, w := range wordDict {
        word[w] = true
    }

    dp := make([]bool, n+1)

    dp[0] = true

    for end := 1; end <= n; end++ {
        for start := 0; start < end; start++ {
            w := s[start:end]

            if _, ok := word[w]; ok && dp[start] {
                dp[end] = true
            }
        }
    }

    return dp[n]
}