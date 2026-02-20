func makeLargestSpecial(s string) string {
	var dfs func(sub string) string
	dfs = func(sub string) string {
		if len(sub) == 0 {
			return ""
		}

		parts := make([]string, 0)
		balance := 0
		start := 0

		for i := 0; i < len(sub); i++ {
			if sub[i] == '1' {
				balance++
			} else {
				balance--
			}
			if balance == 0 {
				inner := dfs(sub[start+1 : i])
				parts = append(parts, "1"+inner+"0")
				start = i + 1
			}
		}

		sort.Slice(parts, func(i, j int) bool {
			return parts[i] > parts[j]
		})

		var b strings.Builder
		for _, p := range parts {
			b.WriteString(p)
		}
		return b.String()
	}

	return dfs(s)
}