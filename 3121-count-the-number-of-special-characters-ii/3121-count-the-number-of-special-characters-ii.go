func numberOfSpecialChars(word string) int {
	lastLower := make([]int, 26)
	firstUpper := make([]int, 26)

	for i := 0; i < 26; i++ {
		lastLower[i] = -1
		firstUpper[i] = -1
	}

	for i := 0; i < len(word); i++ {
		ch := word[i]

		if ch >= 'a' && ch <= 'z' {
			lastLower[ch-'a'] = i
		} else {
			if firstUpper[ch-'A'] == -1 {
				firstUpper[ch-'A'] = i
			}
		}
	}

	res := 0
	for i := 0; i < 26; i++ {
		if lastLower[i] != -1 && firstUpper[i] != -1 && lastLower[i] < firstUpper[i] {
			res++
		}
	}

	return res
}