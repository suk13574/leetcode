func longestBalanced(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		seen := make([]int, 26)
		countSet := make(map[int]int)
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'

			if _, ok := countSet[seen[b]]; ok {
				countSet[seen[b]]--
				if countSet[seen[b]] == 0 {
					delete(countSet, seen[b])
				}
			}

			seen[b]++
			countSet[seen[b]]++

			if len(countSet) == 1 {
				res = max(res, j-i+1)
			}

		}
	}

	return res
}