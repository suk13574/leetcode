func minimumArea(grid [][]int) int {
	x1, y1 := 1001, 1001
	x2, y2 := -1, -1

	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			cnt++
			x1 = min(x1, i)
			x2 = max(x2, i)
			y1 = min(y1, j)
			y2 = max(y2, j)
		}
	}

	if cnt == 1 {
		return 1
	}

	x := x2 - x1 + 1
	y := y2 - y1 + 1

	return x * y
}