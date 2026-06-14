func mapWordWeights(words []string, weights []int) string {

	sumWeights := func(s string) int {
		res := 0
		for i := 0; i < len(s); i++ {
			idx := s[i] - 'a'
			res = (res + weights[idx]) % 26
		}
		return res
	}

	res := make([]byte, 0)
	for _, word := range words {
		weight := sumWeights(word)
		ch := byte('z' - weight)
		res = append(res, ch)
	}

	return string(res)
}