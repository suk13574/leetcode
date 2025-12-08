func countTriples(n int) int {
	squre := make([]int, n+1)
	isSqure := make(map[int]struct{}, n)
	for i := 1; i <= n; i++ {
		squre[i] = i * i
		isSqure[squre[i]] = struct{}{}
	}

	res := 0
	for a := 1; a <= n; a++ {
		as := squre[a]
		for b := 1; b <= n; b++ {
			bs := squre[b]

			cs := as + bs
			if _, ok := isSqure[cs]; ok {
				res++
			}
		}
	}

	return res
}