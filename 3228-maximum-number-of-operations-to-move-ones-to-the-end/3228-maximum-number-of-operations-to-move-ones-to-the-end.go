func maxOperations(s string) int {
	var ones, res int
	for i := 0; i+1 < len(s); i++ {
		if s[i] == '1' {
			ones++
			if s[i+1] == '0' {
				res += ones
			}
		}
	}
	return res
}