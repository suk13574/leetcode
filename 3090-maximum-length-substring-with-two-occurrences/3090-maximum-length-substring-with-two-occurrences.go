func maximumLengthSubstring(s string) int {
	count := make([]int, 26)

	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		idx := s[r] - 'a'

		count[idx]++

		for count[idx] > 2 {
			count[s[l]-'a']--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}