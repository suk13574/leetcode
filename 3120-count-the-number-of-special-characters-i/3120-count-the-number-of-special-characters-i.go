func numberOfSpecialChars(word string) int {
	upper := make([]bool, 26)
	lower := make([]bool, 26)

	for i := 0; i < len(word); i++ {
		ch := word[i]

		if ch >= 'a' && ch <= 'z' {
			lower[ch-'a'] = true
		} else {
			upper[ch-'A'] = true
		}
	}

	res := 0
	for i := 0; i < 26; i++ {
		if upper[i] && lower[i] {
			res++
		}
	}

	return res
}