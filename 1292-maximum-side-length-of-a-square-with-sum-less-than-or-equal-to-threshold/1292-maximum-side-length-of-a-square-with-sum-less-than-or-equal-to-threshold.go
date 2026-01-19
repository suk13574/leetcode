func maxSideLength(mat [][]int, threshold int) int {
	m := len(mat)
	n := len(mat[0])

	prefix := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		prefix[i] = make([]int64, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i+1][j] + prefix[i][j+1] - prefix[i][j] + int64(mat[i][j])
		}
	}

	getSum := func(r, c, k int) int64 {
		return prefix[r+k][c+k] - prefix[r][c+k] - prefix[r+k][c] + prefix[r][c]
	}
	can := func(k int) bool {
		for r := 0; r+k <= m; r++ {
			for c := 0; c+k <= n; c++ {
				if getSum(r, c, k) <= int64(threshold) {
					return true
				}
			}
		}
		return false
	}

	lo, hi := 0, min(m, n)
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if can(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}