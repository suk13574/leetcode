func pyramidTransition(bottom string, allowed []string) bool {
	pattern := make(map[string][]byte)
	for _, a := range allowed {
		p := a[:2]
		if _, ok := pattern[p]; !ok {
			pattern[p] = []byte{}
		}
		pattern[p] = append(pattern[p], a[2])
	}

	memo := make(map[string]bool)
	visit := make(map[string]bool)

	var dfs func(string) bool
	dfs = func(row string) bool {
		if len(row) == 1 {
			return true
		}

		if visit[row] {
			return memo[row]
		}
		visit[row] = true

		nextRow := make([]byte, 0, len(row)-1)

		var build func(i int) bool
		build = func(i int) bool {
			if i == len(row)-1 {
				return dfs(string(nextRow))
			}

			pair := row[i : i+2]
			cands, ok := pattern[pair]
			if !ok {
				return false
			}

			for _, c := range cands {
				nextRow = append(nextRow, c)
				if build(i + 1) {
					return true
				}
				nextRow = nextRow[:len(nextRow)-1]
			}

			return false
		}

		res := build(0)
		memo[row] = res
		return res
	}

	return dfs(bottom)
}