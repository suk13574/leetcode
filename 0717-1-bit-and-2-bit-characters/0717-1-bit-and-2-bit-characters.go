func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	for i := 0; i < n; i++ {
		if i == n-1 {
			return true
		}
		if bits[i] == 1 {
			i++
		}
	}

	return false
}