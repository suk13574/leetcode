func findMaxForm(strs []string, m int, n int) int {
	cnt := make([][2]int, len(strs))
	for i, s := range strs {
		z, o := 0, 0
		for j := 0; j < len(s); j++ {
			if s[j] == '0' {
				z++
			} else {
				o++
			}
		}
		cnt[i] = [2]int{z, o}
	}

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for _, c := range cnt {
		z, o := c[0], c[1]

		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				if v := dp[i-z][j-o] + 1; v > dp[i][j] {
					dp[i][j] = v
				}
			}
		}
	}
	return dp[m][n]
}