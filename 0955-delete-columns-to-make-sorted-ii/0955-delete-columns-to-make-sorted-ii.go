func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	sorted := make([]bool, n-1)
	res := 0

	for col := 0; col < m; col++ {
		bad := false
		for i := 0; i < n-1; i++ {
			if sorted[i] {
				continue
			}

			if strs[i][col] > strs[i+1][col] {
				bad = true
				break
			}
		}

		if bad {
			res++
			continue
		}

		for i := 0; i < n-1; i++ {
			if sorted[i] {
				continue
			}
			if strs[i][col] < strs[i+1][col] {
				sorted[i] = true
			}
		}

	}
	return res
}