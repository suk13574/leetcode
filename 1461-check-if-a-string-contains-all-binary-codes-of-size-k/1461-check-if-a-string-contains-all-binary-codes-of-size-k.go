func hasAllCodes(s string, k int) bool {
	set := make(map[string]struct{})
	for i := 0; i+k <= len(s); i++ {
		subs := s[i : i+k]
		if _, ok := set[subs]; !ok {
			set[subs] = struct{}{}
		}
	}

	if len(set) == (1<<k) {
		return true
	}
	return false
}