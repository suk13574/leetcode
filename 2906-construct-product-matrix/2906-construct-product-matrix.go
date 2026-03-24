func constructProductMatrix(grid [][]int) [][]int {
	const MOD = 12345

	r, c := len(grid), len(grid[0])
	n := r * c
	prefix := make([]int, n+1)
	suffix := make([]int, n+1)

	prefix[0] = 1
	for k := 0; k < n; k++ {
		i, j := (k)/c, (k)%c
		prefix[k+1] = (prefix[k] * grid[i][j]) % MOD
	}

	suffix[n] = 1
	for k := n - 1; k >= 0; k-- {
		i, j := (k)/c, (k)%c
		suffix[k] = (suffix[k+1] * grid[i][j]) % MOD
	}

	p := make([][]int, r)
	for i := 0; i < r; i++ {
		p[i] = make([]int, c)
	}

	for k := 0; k < n; k++ {
		i, j := k/c, k%c
		p[i][j] = (prefix[k] * suffix[k+1]) % MOD
	}

	return p
}
