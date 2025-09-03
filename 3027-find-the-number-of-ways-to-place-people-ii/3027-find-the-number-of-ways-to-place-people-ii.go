func compress(nums []int) map[int]int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)

	comp := make(map[int]int)
	idx := 1
	for _, num := range copyNums {
		if _, ok := comp[num]; ok {
			continue
		}

		comp[num] = idx
		idx++
	}
	return comp
}

func numberOfPairs(points [][]int) int {
	n := len(points)
	xNums, yNums := make([]int, n), make([]int, n)
	for i, p := range points {
		xNums[i] = p[0]
		yNums[i] = p[1]
	}

	cx := compress(xNums)
	cy := compress(yNums)

	cxArr, cyArr := make([]int, n), make([]int, n)
	for i, p := range points {
		cxArr[i] = cx[p[0]]
		cyArr[i] = cy[p[1]]
	}

	nx := len(cx) + 1
	ny := len(cy) + 1
	grid := make([][]int, nx)
	for i := 0; i < nx; i++ {
		grid[i] = make([]int, ny)
	}

	for i := 0; i < n; i++ {
		grid[cxArr[i]][cyArr[i]]++
	}

	prefix := make([][]int, nx)
	for i := 0; i < nx; i++ {
		prefix[i] = make([]int, ny)
	}

	for i := 1; i < nx; i++ {
		row := 0
		for j := 1; j < ny; j++ {
			row += grid[i][j]
			prefix[i][j] = prefix[i-1][j] + row
		}
	}

	sumRect := func(x1, y1, x2, y2 int) int {
		return prefix[x2][y2] - prefix[x1-1][y2] - prefix[x2][y1-1] + prefix[x1-1][y1-1]
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := cxArr[i], cyArr[i]
			xj, yj := cxArr[j], cyArr[j]

			if xi <= xj && yj <= xj {
				if sumRect(xi, yj, xj, yi) == 2 {
					result++
				}
			}

			if xj <= xi && yi <= yj {
				if sumRect(xj, yi, xi, yj) == 2 {
					result++
				}
			}
		}
	}

	return result
}