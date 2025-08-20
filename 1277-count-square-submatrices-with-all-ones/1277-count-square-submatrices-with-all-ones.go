func countSquares(matrix [][]int) int {
	result := 0
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}

			if i-1 >= 0 && j-1 >= 0 {
				matrix[i][j] += min(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
			}
			result += matrix[i][j]
		}
	}
	return result
}