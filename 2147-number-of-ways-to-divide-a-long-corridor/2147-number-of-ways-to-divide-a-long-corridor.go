func numberOfWays(corridor string) int {
	const MOD int64 = 1000000007
	n := len(corridor)

	totalS := 0
	for i := 0; i < n; i++ {
		if corridor[i] == 'S' {
			totalS++
		}
	}
	if totalS == 0 || totalS%2 == 1 {
		return 0
	}

	var res int64 = 1
	countS := 0
	gapP := 0
	for i := 0; i < n; i++ {
		c := corridor[i]

		if c == 'S' {
			countS++
			if countS%2 == 1 && countS > 1 {
				res = (res * int64(gapP+1)) % MOD
				gapP = 0
			}
		} else {
			if countS > 0 && countS%2 == 0 && countS < totalS {
				gapP++
			}
		}
	}

	return int(res)
}