func canBeEqual(s1 string, s2 string) bool {
	b := []byte(s2)

	candidates := []string{
		string(b),
		string([]byte{b[2], b[1], b[0], b[3]}),
		string([]byte{b[0], b[3], b[2], b[1]}),
		string([]byte{b[2], b[3], b[0], b[1]}),
	}

	for _, c := range candidates {
		if s1 == c {
			return true
		}
	}

	return false
}