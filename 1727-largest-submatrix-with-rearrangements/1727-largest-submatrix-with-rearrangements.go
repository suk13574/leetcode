func largestSubmatrix(matrix [][]int) int {
	r, c := len(matrix), len(matrix[0])

	for j := 0; j < c; j++ {
		for i := 1; i < r; i++ {
			if matrix[i][j] == 0 {
				continue
			}

			matrix[i][j] += matrix[i-1][j]
		}
	}

	res := 0
	rows := make([]int, c)
	for i := 0; i < r; i++ {
		copy(rows, matrix[i])
		sort.Ints(rows)

		for j := 0; j < c; j++ {
			w := c - j
			h := rows[j]
			res = max(res, w*h)
		}
	}

	return res
}