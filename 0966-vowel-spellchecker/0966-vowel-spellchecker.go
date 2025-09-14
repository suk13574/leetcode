func UnifyVowels(s string) string {
	ss := []byte(s)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			ss[i] = '_'
		}
	}
	return string(ss)
}

func spellchecker(wordlist []string, queries []string) []string {
	str := make(map[string]string)
	strLower := make(map[string]string)
	strLowerVowel := make(map[string]string)

	for _, w := range wordlist {
		str[w] = w

		wLower := strings.ToLower(w)
		if _, ok := strLower[wLower]; !ok {
			strLower[wLower] = w
		}

		wLowerVowel := UnifyVowels(wLower)
		if _, ok := strLowerVowel[wLowerVowel]; !ok {
			strLowerVowel[wLowerVowel] = w
		}
	}

	answer := make([]string, len(queries))
	for i := 0; i < len(queries); i++ {
		q := queries[i]

		if s, ok := str[q]; ok {
			answer[i] = s
			continue
		}

		q = strings.ToLower(q)
		if s, ok := strLower[q]; ok {
			answer[i] = s
			continue
		}

		q = UnifyVowels(q)
		if s, ok := strLowerVowel[q]; ok {
			answer[i] = s
			continue
		}

		answer[i] = ""
	}

	return answer
}