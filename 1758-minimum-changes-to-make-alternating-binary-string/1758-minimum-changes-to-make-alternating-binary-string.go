func minOperations(s string) int {
	diff := 0

	for i := 0; i < len(s); i++ {
		expected := byte('0')
		if i&1 == 1 {
			expected = byte('1')
		}

		if s[i] != expected {
			diff++
		}
	}

	return min(diff, len(s)-diff)

}