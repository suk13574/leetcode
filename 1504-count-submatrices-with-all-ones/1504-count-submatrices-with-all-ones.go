func numSubmat(mat [][]int) int {
	result := 0

	m := len(mat)
	n := len(mat[0])

	width := make([][]int, m)
	for i := 0; i < m; i++ {
		width[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}

			width[i][j] = mat[i][j]
			if j-1 >= 0 {
				width[i][j] += width[i][j-1]
			}

			minWidth := width[i][j]
			for k := i; k >= 0; k-- {
				minWidth = min(minWidth, width[k][j])
				if minWidth == 0 {
					continue
				}
				result += minWidth
			}
		}
	}
	return result
}