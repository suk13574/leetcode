func zigZagArrays(n int, l int, r int) int {
	const MOD = 1_000_000_007

	m := r - l + 1
	up := make([]int64, m+2)
	down := make([]int64, m+2)

	for i := 1; i <= m; i++ {
		up[i] = int64(i - 1)
		down[i] = int64(m - i)
	}

	for length := 3; length <= n; length++ {
		newUp := make([]int64, m+2)
		newDown := make([]int64, m+2)

		prefix := int64(0)
		for j := 1; j <= m; j++ {
			newUp[j] = prefix
			prefix = (prefix + down[j]) % MOD
		}

		suffix := int64(0)
		for j := m; j >= 1; j-- {
			newDown[j] = suffix
			suffix = (suffix + up[j]) % MOD
		}

		up = newUp
		down = newDown
	}

	res := int64(0)
	for i := 1; i <= m; i++ {
		res = (res + up[i] + down[i]) % MOD
	}

	return int(res)
}