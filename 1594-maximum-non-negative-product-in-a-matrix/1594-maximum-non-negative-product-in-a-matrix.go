func maxProductPath(grid [][]int) int {
	const MOD = 1_000_000_007

	r, c := len(grid), len(grid[0])
	maxDp := make([][]int64, r)
	minDp := make([][]int64, r)
	for i := 0; i < r; i++ {
		maxDp[i] = make([]int64, c)
		minDp[i] = make([]int64, c)
	}

	maxDp[0][0] = int64(grid[0][0])
	minDp[0][0] = int64(grid[0][0])

	for i := 1; i < r; i++ {
		v := int64(grid[i][0])
		maxDp[i][0] = maxDp[i-1][0] * v
		minDp[i][0] = minDp[i-1][0] * v
	}

	for j := 1; j < c; j++ {
		v := int64(grid[0][j])
		maxDp[0][j] = maxDp[0][j-1] * v
		minDp[0][j] = minDp[0][j-1] * v
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			v := int64(grid[i][j])

			left1 := maxDp[i-1][j] * v
			left2 := minDp[i-1][j] * v
			up1 := maxDp[i][j-1] * v
			up2 := minDp[i][j-1] * v

			maxDp[i][j] = max(left1, max(left2, max(up1, up2)))
			minDp[i][j] = min(left1, min(left2, min(up1, up2)))
		}
	}
	if maxDp[r-1][c-1] < 0 {
		return -1
	}
	return int(maxDp[r-1][c-1] % MOD)
}