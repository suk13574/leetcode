func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	n := len(s)
	for i := 0; i < n; i++ {
		if s == goal {
			return true
		}

		s = s[1:] + s[:1]
	}

	return false
}