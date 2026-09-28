func reverseParentheses(s string) string {
	idx := 0

	reverse := func(str []byte) string {
		for l, r := 0, len(str)-1; l < r; l, r = l+1, r-1 {
			str[l], str[r] = str[r], str[l]
		}

		return string(str)
	}

	var dfs func() string

	dfs = func() string {
		sub := []byte{}

		for idx < len(s) {
			ch := s[idx]
			idx++

			if ch == '(' {
				inner := dfs()
				sub = append(sub, inner...)
			} else if ch == ')' {
				return reverse(sub)
			} else {
				sub = append(sub, ch)
			}
		}
		return string(sub)
	}

	return dfs()
}