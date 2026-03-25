func canPartitionGrid(grid [][]int) bool {
	r, c := len(grid), len(grid[0])

	hsum := make([]int64, r)
	vsum := make([]int64, c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v := int64(grid[i][j])

			hsum[i] += v
			vsum[j] += v
		}
	}

	for i := 1; i < r; i++ {
		hsum[i] += hsum[i-1]
	}
	for j := 1; j < c; j++ {
		vsum[j] += vsum[j-1]
	}

	total := hsum[r-1]
	if total&1 == 1 {
		return false
	}
	target := total / 2

	for i := 0; i < r; i++ {
		if target == hsum[i] {
			return true
		}

		if target < hsum[i] {
			break
		}
	}

	for j := 0; j < c; j++ {
		if target == vsum[j] {
			return true
		}

		if target < vsum[j] {
			break
		}
	}

	return false
}