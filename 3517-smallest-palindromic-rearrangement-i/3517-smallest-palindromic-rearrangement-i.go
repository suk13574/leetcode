func smallestPalindrome(s string) string {
	n := len(s)

	alphabet := make([]int, 26)
	for i := 0; i < n; i++ {
		alphabet[s[i]-'a']++
	}

	res := make([]byte, n)
	idx := 0
	for i := 0; i < 26; i++ {
		ch := byte(i) + 'a'
		for alphabet[i] > 1 {
			res[idx] = ch
			res[n-1-idx] = ch
			idx++

			alphabet[i] -= 2
		}

		if alphabet[i] == 1 {
			res[n/2] = ch
		}
	}

	return string(res)
}