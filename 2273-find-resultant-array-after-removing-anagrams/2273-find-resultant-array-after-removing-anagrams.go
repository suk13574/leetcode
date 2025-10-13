func isAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	b1, b2 := []byte(w1), []byte(w2)
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	return string(b1) == string(b2)
}

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	for i := 1; i < len(words); i++ {
		w1 := res[len(res)-1]
		w2 := words[i]

		if isAnagram(w1, w2) {
			continue
		}
		res = append(res, w2)
	}

	return res
}