func separateSquares(squares [][]int) float64 {
	lo := 1e30
	hi := -1e30

	total := 0.0
	for _, s := range squares {
		y := float64(s[1])
		l := float64(s[2])
		lo = min(lo, y)
		hi = max(hi, y+l)
		total += l * l
	}

	target := total / 2.0

	belowArea := func(y float64) float64 {
		sum := 0.0
		for _, s := range squares {
			yi := float64(s[1])
			l := float64(s[2])
			top := yi + l

			if y <= yi {
				continue
			} else if y >= top {
				sum += l * l
			} else {
				sum += l * (y - yi)
			}
		}
		return sum
	}

	eps := 1e-7
	for (hi - lo) > eps {
		mid := (lo + hi) / 2.0
		if belowArea(mid) < target {
			lo = mid
		} else {
			hi = mid
		}
	}
	return hi
}