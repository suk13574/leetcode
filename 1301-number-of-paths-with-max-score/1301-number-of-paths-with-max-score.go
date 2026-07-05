func pathsWithMaxScore(board []string) []int {
	const MOD int64 = 1_000_000_007

	m, n := len(board), len(board[0])

	dpScore := make([][]int, m)
	dpPath := make([][]int64, m)
	for i := 0; i < m; i++ {
		dpScore[i] = make([]int, n)
		dpPath[i] = make([]int64, n)

		for j := 0; j < n; j++ {
			dpScore[i][j] = -1
		}
	}

	moves := [][]int{{0, 1}, {1, 0}, {1, 1}}

	dpScore[0][0] = 0
	dpPath[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			if dpPath[i][j] == 0 {
				continue
			}

			for _, move := range moves {
				nx, ny := i+move[0], j+move[1]

				if nx >= m || ny >= n {
					continue
				}

				if board[nx][ny] == 'X' {
					continue
				}

				add := 0
				if board[nx][ny] >= '1' && board[nx][ny] <= '9' {
					add += int(board[nx][ny] - '0')
				}
				nextScore := dpScore[i][j] + add

				if nextScore > dpScore[nx][ny] {
					dpScore[nx][ny] = nextScore
					dpPath[nx][ny] = dpPath[i][j]
				} else if nextScore == dpScore[nx][ny] {
					dpPath[nx][ny] = (dpPath[i][j] + dpPath[nx][ny]) % MOD
				}
			}
		}
	}

	if dpPath[m-1][n-1] == 0 {
		return []int{0, 0}
	}

	return []int{dpScore[m-1][n-1], int(dpPath[m-1][n-1])}
}