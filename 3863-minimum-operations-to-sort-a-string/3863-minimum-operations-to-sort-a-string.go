func minOperations(s string) int {
	n := len(s)

	sorted := true
	minChar := s[0]
	maxChar := s[0]
	minCnt := 1
	maxCnt := 1

	for i := 1; i < n; i++ {
		if s[i-1] > s[i] {
			sorted = false
		}

		if s[i] < minChar {
			minChar = s[i]
			minCnt = 0
		}

		if s[i] > maxChar {
			maxChar = s[i]
			maxCnt = 0
		}

		if s[i] == minChar {
			minCnt++
		}

		if s[i] == maxChar {
			maxCnt++
		}
	}

	if sorted {
		return 0
	}

	if n == 2 {
		return -1
	}

	if s[0] == minChar || s[n-1] == maxChar {
		return 1
	}

	if s[0] == maxChar && maxCnt == 1 &&
		s[n-1] == minChar && minCnt == 1 {
		return 3
	}

	return 2
}