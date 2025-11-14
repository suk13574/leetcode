func rangeAddQueries(n int, queries [][]int) [][]int {
	diff := make([][]int, n+1)
	for i := 0; i < len(diff); i++ {
		diff[i] = make([]int, n+1)
	}

	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		diff[r1][c1] += 1
		diff[r1][c2+1] -= 1
		diff[r2+1][c1] -= 1
		diff[r2+1][c2+1] += 1
	}

	var preSum int
	sum := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		sum[i] = make([]int, n+1)
		preSum = 0
		for j := 0; j < n+1; j++ {
			preSum += diff[i][j]
			sum[i][j] = preSum
		}
	}

	for i := 0; i < n+1; i++ {
		preSum = 0
		for j := 0; j < n+1; j++ {
			sum[j][i] += preSum
			preSum = sum[j][i]

		}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = sum[i][:n]
	}

	return res
}