import "sort"

type TeleCell struct {
	val  int
	i, j int
}

func minCost(grid [][]int, k int) int {
	const INF = 1 << 60
	m := len(grid)
	n := len(grid[0])

	teleMap := make([]TeleCell, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			teleMap = append(teleMap, TeleCell{val: grid[i][j], i: i, j: j})
		}
	}

	sort.Slice(teleMap, func(a, b int) bool {
		if teleMap[a].val != teleMap[b].val {
			return teleMap[a].val < teleMap[b].val
		}
		if teleMap[a].i != teleMap[b].i {
			return teleMap[a].i < teleMap[b].i
		}
		return teleMap[a].j < teleMap[b].j
	})

	dp := make([][]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	normalMove := func() {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				best := dp[i][j]

				if i > 0 && dp[i-1][j] < best {
					best = dp[i-1][j]
				}
				if j > 0 && dp[i][j-1] < best {
					best = dp[i][j-1]
				}

				if best != INF {
					dp[i][j] = min(dp[i][j], best+int64(grid[i][j]))
				}
			}
		}
	}

	teleMove := func() {
		tn := len(teleMap)

		var groupMin int64
		groupMin = INF
		currVal := -1
		for idx := 0; idx < tn; idx++ {
			v, i, j := teleMap[idx].val, teleMap[idx].i, teleMap[idx].j
			if v != currVal {
				currVal = v
				groupMin = dp[i][j]
			} else {
				if dp[i][j] < groupMin {
					groupMin = dp[i][j]
				} else {
					dp[i][j] = groupMin
				}
			}
		}

		subffixMin := dp[teleMap[tn-1].i][teleMap[tn-1].j]
		for idx := tn - 1; idx >= 0; idx-- {
			i, j := teleMap[idx].i, teleMap[idx].j
			if dp[i][j] < subffixMin {
				subffixMin = dp[i][j]
			} else {
				dp[i][j] = subffixMin
			}
		}
	}

	for t := k; t >= 0; t-- {
		normalMove()
		if t > 0 {
			teleMove()
		}
	}

	return int(dp[m-1][n-1])
}