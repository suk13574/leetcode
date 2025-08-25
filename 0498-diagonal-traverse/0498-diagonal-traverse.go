func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	up := true
	x, y := 0, 0

	result := make([]int, m*n)

	for i := 0; i < m*n; i++ {
		result[i] = mat[x][y]
		if up {

			if y == n-1 {
				x++
				up = false
			} else if x == 0 {
				y++
				up = false
			} else {
				x--
				y++
			}

		} else {
			if x == m-1 {
				y++
				up = true
			} else if y == 0 {
				x++
				up = true
			} else {
				x++
				y--
			}
		}
	}

	return result
}