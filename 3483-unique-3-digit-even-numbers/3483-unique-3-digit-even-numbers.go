func totalNumbers(digits []int) int {
	cnt := make([]int, 10)
	for _, d := range digits {
		cnt[d]++
	}

	var dfs func(int, int) int
	dfs = func(pos, k int) int {
		if pos == k {
			return 1
		}

		res := 0

		for d := 0; d <= 9; d++ {
			if cnt[d] == 0 {
				continue
			}

			if pos == 0 && d == 0 {
				continue
			}

			if pos == k-1 && d%2 != 0 {
				continue
			}

			cnt[d]--
			res += dfs(pos+1, 3)
			cnt[d]++
		}

		return res
	}

	return dfs(0, 3)
}