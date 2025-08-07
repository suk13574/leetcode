func maxCollectedFruits(fruits [][]int) int {
	result := 0
	n := len(fruits)

	firstChild := 0
	for i := 0; i < n; i++ {
		firstChild += fruits[i][i]
	}

	result = firstChild + secondChild(fruits) + thirdCHild(fruits) - 2*fruits[n-1][n-1]

	return result
}

func secondChild(fruits [][]int) int {
	n := len(fruits)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][n-1] = fruits[0][n-1]
	moves := [][]int{{1, -1}, {1, 0}, {1, 1}}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j && i != n-1 && j != n-1 {
				continue
			}

			for _, move := range moves {
				ni := i - move[0]
				nj := j - move[1]

				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if ni < nj && nj < n-1-ni {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[ni][nj]+fruits[i][j])
			}

		}
	}
	return dp[n-1][n-1]
}

func thirdCHild(fruits [][]int) int {
	n := len(fruits)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[n-1][0] = fruits[n-1][0]
	moves := [][]int{{-1, 1}, {0, 1}, {1, 1}}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if i <= j && j != n-1 && i != n-1 {
				continue
			}

			for _, move := range moves {
				ni := i - move[0]
				nj := j - move[1]

				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}
				if nj < ni && ni < n-1-nj {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[ni][nj]+fruits[i][j])
			}
		}
	}

	return dp[n-1][n-1]
}