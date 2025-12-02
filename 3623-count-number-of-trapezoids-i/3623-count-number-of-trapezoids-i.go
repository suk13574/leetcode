func countTrapezoids(points [][]int) int {
	const MOD = 1000000007

	yMap := make(map[int]int64)
	for _, p := range points {
		y := p[1]
		yMap[y]++
	}

	var res int64
	var sum int64
	for _, n := range yMap {
		if n < 2 {
			continue
		}

		hCnt := (n * (n - 1) / 2) % MOD

		res = (res + sum*hCnt) % MOD
		sum = (sum + hCnt) % MOD

	}

	return int(res)
}