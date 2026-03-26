func canPartitionGrid(grid [][]int) bool {
	r, c := len(grid), len(grid[0])

	hsum := make([]int64, r)
	vsum := make([]int64, c)
	hash := make(map[int][][]int)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v := int64(grid[i][j])

			hsum[i] += v
			vsum[j] += v
			hash[grid[i][j]] = append(hash[grid[i][j]], []int{i, j})
		}
	}

	for i := 1; i < r; i++ {
		hsum[i] += hsum[i-1]
	}
	for j := 1; j < c; j++ {
		vsum[j] += vsum[j-1]
	}

	total := hsum[r-1]

	isRemoveOk := func(x1, y1, x2, y2 int, gap int64) bool {
		coordinates, ok := hash[int(gap)]
		if !ok {
			return false
		}

		for _, p := range coordinates {
			x, y := p[0], p[1]
			if !(x1 <= x && x <= x2 && y1 <= y && y <= y2) {
				continue
			}

			if x1 == x2 && y != y1 && y != y2 {
				continue
			}

			if y1 == y2 && x != x1 && x != x2 {
				continue
			}
			return true
		}
		return false
	}

	// horizontal
	for i := 0; i < r-1; i++ {
		top := hsum[i]
		bottom := total - top

		if top == bottom {
			return true
		}

		if top > bottom {
			if isRemoveOk(0, 0, i, c-1, top-bottom) {
				return true
			}
		} else {
			if isRemoveOk(i+1, 0, r-1, c-1, bottom-top) {
				return true
			}
		}
	}

	// vertical
	for j := 0; j < c-1; j++ {
		left := vsum[j]
		right := total - left

		if left == right {
			return true
		}

		if left > right {
			if isRemoveOk(0, 0, r-1, j, left-right) {
				return true
			}
		} else {
			if isRemoveOk(0, j+1, r-1, c-1, right-left) {
				return true
			}
		}
	}

	return false
}