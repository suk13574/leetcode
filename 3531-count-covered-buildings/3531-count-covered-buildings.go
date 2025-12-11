type MinMax struct {
	min int
	max int
}

func countCoveredBuildings(n int, buildings [][]int) int {
	xLineY := make(map[int]*MinMax)
	yLineX := make(map[int]*MinMax)

	for _, b := range buildings {
		x, y := b[0], b[1]
		if _, ok := xLineY[x]; !ok {
			xLineY[x] = &MinMax{y, y}
		}
		xLineY[x].max = max(xLineY[x].max, y)
		xLineY[x].min = min(xLineY[x].min, y)

		if _, ok := yLineX[y]; !ok {
			yLineX[y] = &MinMax{x, x}
		}
		yLineX[y].max = max(yLineX[y].max, x)
		yLineX[y].min = min(yLineX[y].min, x)

	}

	res := 0
	for _, b := range buildings {
		x, y := b[0], b[1]
		if xLineY[x].min < y && y < xLineY[x].max {
			if yLineX[y].min < x && x < yLineX[y].max {
				res++
			}
		}

	}
	return res

}