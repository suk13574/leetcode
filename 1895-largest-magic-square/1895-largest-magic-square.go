func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	rowSums := make([][]int64, m)
	for i := 0; i < m; i++ {
		rowSums[i] = make([]int64, n+1)
	}

	colSums := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		colSums[i] = make([]int64, n)
	}

	diagSums1 := make([][]int64, m+1)
	diagSums2 := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		diagSums1[i] = make([]int64, n+1)

		diagSums2[i] = make([]int64, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := int64(grid[i][j])
			rowSums[i][j+1] = rowSums[i][j] + v
			colSums[i+1][j] = colSums[i][j] + v
			diagSums1[i+1][j+1] = diagSums1[i][j] + v
			diagSums2[i+1][j] = diagSums2[i][j+1] + v
		}
	}

	getRow := func(r, c1, c2 int) int64 {
		return rowSums[r][c2] - rowSums[r][c1]
	}
	getCol := func(r1, r2, c int) int64 {
		return colSums[r2][c] - colSums[r1][c]
	}
	getD1 := func(r, c, k int) int64 {
		return diagSums1[r+k][c+k] - diagSums1[r][c]
	}
	getD2 := func(r, c, k int) int64 {
		return diagSums2[r+k][c] - diagSums2[r][c+k]
	}

	maxK := min(m, n)
	for k := maxK; k >= 2; k-- {
		for r := 0; r+k <= m; r++ {
			for c := 0; c+k <= n; c++ {
				target := getD1(r, c, k)
				if getD2(r, c, k) != target {
					continue
				}

				ok := true
				for i := 0; i < k && ok; i++ {
					if getRow(r+i, c, c+k) != target {
						ok = false
					}
				}
				for j := 0; j < k && ok; j++ {
					if getCol(r, r+k, c+j) != target {
						ok = false
					}
				}

				if ok {
					return k
				}
			}
		}
	}
	return 1

}