func countAlphabet(w string) [26]byte {
	count := [26]byte{}
	for i := 0; i < len(w); i++ {
		count[w[i]-'a']++
	}
	return count
}

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	cntW1 := countAlphabet(res[0])
	for i := 1; i < len(words); i++ {
		w2 := words[i]
		cntW2 := countAlphabet(w2)

		if cntW1 != cntW2 {
			res = append(res, w2)
			cntW1 = cntW2
		}
	}

	return res
}