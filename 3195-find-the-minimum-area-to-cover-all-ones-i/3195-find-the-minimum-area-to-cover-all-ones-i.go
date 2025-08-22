import "math"

func minimumArea(grid [][]int) int {
	x1, y1 := math.MaxInt, math.MaxInt
	x2, y2 := math.MinInt, math.MinInt
	found := false

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if !found {
					x1, y1, x2, y2 = i, j, i, j
					found = true
				} else {
					x1 = min(x1, i)
					x2 = max(x2, i)
					y1 = min(y1, j)
					y2 = max(y2, j)
				}
			}

		}
	}

	if !found {
		return 0
	}

	x := x2 - x1 + 1
	y := y2 - y1 + 1

	return x * y
}