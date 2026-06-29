func numOfStrings(patterns []string, word string) int {
	res := 0

	for _, p := range patterns {
		if isSubstring(p, word) {
			res++
		}
	}

	return res
}

func isSubstring(pattern, word string) bool {
	lps := buildLPS(pattern)

	i := 0 // word index
	j := 0 // pattern index

	for i < len(word) {
		if word[i] == pattern[j] {
			i++
			j++

			if j == len(pattern) {
				return true
			}
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return false
}

func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))

	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length > 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}