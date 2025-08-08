func soupServings(n int) float64 {
	if n > 4800 {
		return 1
	}

	dp := make(map[[2]int]float64)

	var probabailySoup func(a, b int) float64

	N := (n + 24) / 25

	probabailySoup = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 && 0 < b {
			return 1
		}
		if 0 < a && b <= 0 {
			return 0
		}

		key := [2]int{a, b}
		if val, ok := dp[key]; ok {
			return val
		}

		case1 := probabailySoup(a-4, b)
		case2 := probabailySoup(a-1, b-3)
		case3 := probabailySoup(a-2, b-2)
		case4 := probabailySoup(a-3, b-1)
		dp[key] = 0.25 * (case1 + case2 + case3 + case4)

		return dp[key]
	}

	answer := probabailySoup(N, N)

	return answer

}