func sortVowels(s string) string {
	isVowel := func(b rune) bool {
		switch b {
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	idxs := make([]int, 0)
	vowels := make([]rune, 0)
	for i, b := range s {
		if isVowel(b) {
			idxs = append(idxs, i)
			vowels = append(vowels, b)
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	t := []rune(s)
	for i, v := range vowels {
		t[idxs[i]] = v
	}

	return string(t)
}