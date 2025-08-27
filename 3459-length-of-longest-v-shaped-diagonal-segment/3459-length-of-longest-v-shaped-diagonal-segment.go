func lenOfVDiagonal(grid [][]int) int {
	result := 0
	m, n := len(grid), len(grid[0])
	dirs := [4][2]int{{-1, +1}, {+1, +1}, {+1, -1}, {-1, -1}}

	straight := make([][][]int, 4)
	run2 := make([][][]int, 4)
	run0 := make([][][]int, 4)

	for d := 0; d < 4; d++ {
		straight[d] = make([][]int, m)
		run2[d] = make([][]int, m)
		run0[d] = make([][]int, m)
		for i := 0; i < m; i++ {
			straight[d][i] = make([]int, n)
			run2[d][i] = make([]int, n)
			run0[d][i] = make([]int, n)
		}
	}

	for d := 0; d < 4; d++ {
		di, dj := dirs[d][0], dirs[d][1]

		iStart, iEnd, iStep := 0, m, 1
		jStart, jEnd, jStep := 0, n, 1
		if di == -1 {
			iStart, iEnd, iStep = m-1, -1, -1
		}
		if dj == -1 {
			jStart, jEnd, jStep = n-1, -1, -1
		}

		for i := iStart; i != iEnd; i += iStep {
			for j := jStart; j != jEnd; j += jStep {
				if grid[i][j] == 1 {
					straight[d][i][j] = 1
				}

				pi, pj := i-di, j-dj
				if 0 <= pi && pi < m && 0 <= pj && pj < n {
					prevLen := straight[d][pi][pj]

					if prevLen > 0 {
						need := 2
						if prevLen%2 == 0 {
							need = 0
						}
						if grid[i][j] == need {
							straight[d][i][j] = max(straight[d][i][j], prevLen+1)
						}
					}
				}
			}
		}
	}

	for d := 0; d < 4; d++ {
		di, dj := dirs[d][0], dirs[d][1]

		var iStart, iEnd, iStep = 0, m, 1
		var jStart, jEnd, jStep = 0, n, 1
		if di == 1 {
			iStart, iEnd, iStep = m-1, -1, -1
		} 
		if dj == 1 {
			jStart, jEnd, jStep = n-1, -1, -1
		}

		for i := iStart; i != iEnd; i += iStep {
			for j := jStart; j != jEnd; j += jStep {
				ni, nj := i+di, j+dj
				if 0 <= ni && ni < m && 0 <= nj && nj < n {
					if grid[ni][nj] == 2 {
						run2[d][i][j] = 1 + run0[d][ni][nj]
					}
					if grid[ni][nj] == 0 {
						run0[d][i][j] = 1 + run2[d][ni][nj]
					}
				}

			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for d := 0; d < 4; d++ {
				strLen := straight[d][i][j]
				if strLen == 0 {
					continue
				}

				turnD := (d + 1) % 4
				need := 2
				if strLen%2 == 0 {
					need = 0
				}

				count := 0
				if need == 2 {
					count = run2[turnD][i][j]
				} else {
					count = run0[turnD][i][j]
				}
				result = max(result, strLen+count)

			}
		}
	}

	return result
}