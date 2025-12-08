func countTriples(n int) int {
	square := make([]int, n+1)
	isSquare := make(map[int]struct{}, n)
	for i := 1; i <= n; i++ {
		square[i] = i * i
		isSquare[square[i]] = struct{}{}
	}

	res := 0
	for a := 1; a <= n; a++ {
		as := square[a]
		for b := 1; b <= n; b++ {
			bs := square[b]

			cs := as + bs
			if _, ok := isSquare[cs]; ok {
				res++
			}
		}
	}

	return res
}