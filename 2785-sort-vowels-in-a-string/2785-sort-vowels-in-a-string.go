func sortVowels(s string) string {
	isVowel := func(b byte) bool {
		switch b {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	vowels := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	t := []byte(s)
	idx := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			t[i] = vowels[idx]
			idx++
		}
	}

	return string(t)
}