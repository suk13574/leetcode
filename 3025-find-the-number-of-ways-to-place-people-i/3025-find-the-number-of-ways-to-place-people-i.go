func numberOfPairs(points [][]int) int {
	const MAX = 51
	n := len(points)

	grid := make([][]int, MAX+1)
	for i := range grid {
		grid[i] = make([]int, MAX+1)
	}

	for _, p := range points {
		x, y := p[0]+1, p[1]+1
		grid[x][y] = 1
	}

	pref := make([][]int, MAX+1)
	for i := range pref {
		pref[i] = make([]int, MAX+1)
	}
	for x := 1; x <= MAX; x++ {
		row := 0
		for y := 1; y <= MAX; y++ {
			row += grid[x][y]
			pref[x][y] = pref[x-1][y] + row
		}
	}

	sumRect := func(x1, y1, x2, y2 int) int {
		return pref[x2][y2] - pref[x1-1][y2] - pref[x2][y1-1] + pref[x1-1][y1-1]
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			Ax, Ay := points[i][0], points[i][1]
			Bx, By := points[j][0], points[j][1]
			if Ax <= Bx && By <= Ay {
				x1 := min(Ax, Bx) + 1
				y1 := min(Ay, By) + 1
				x2 := max(Ax, Bx) + 1
				y2 := max(Ay, By) + 1
				if sumRect(x1, y1, x2, y2) == 2 {
					result++
				}
			}

		}
	}
	return result
}