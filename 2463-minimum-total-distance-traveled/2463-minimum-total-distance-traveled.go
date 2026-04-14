import "sort"

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })

	r, f := len(robot), len(factory)
	const INF = int64(1e18)

	dp := make([][]int64, r+1)
	for i := 0; i <= r; i++ {
		dp[i] = make([]int64, f+1)
	}

	for i := 0; i < r; i++ {
		dp[i][f] = INF
	}

	for j := f - 1; j >= 0; j-- {
		for i := r; i >= 0; i-- {
			dp[i][j] = dp[i][j+1]

			cost := int64(0)
			for k := 1; k <= factory[j][1] && i+k <= r; k++ {
				cost += abs64(int64(robot[i+k-1]) - int64(factory[j][0]))
				dp[i][j] = min(dp[i][j], cost+dp[i+k][j+1])
			}
		}
	}

	return dp[0][0]
}

func abs64(x int64) int64 {
    if x < 0  {
        return -x
    }
    return x
}