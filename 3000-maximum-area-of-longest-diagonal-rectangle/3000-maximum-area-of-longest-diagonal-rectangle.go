func areaOfMaxDiagonal(dimensions [][]int) int {
	result := 0
	maxDiagonal := 0

	for _, d := range dimensions {
		x, y := d[0], d[1]
		pita := (x * x) + (y * y)

		if maxDiagonal < pita {
			maxDiagonal = pita
			result = x * y
		} else if maxDiagonal == pita {
			result = max(result, x*y)
		}
	}

	return result
}