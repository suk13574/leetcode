func sumZero(n int) []int {
	result := make([]int, n)
	for i := 0; i < n/2; i++ {
		result[i] = -(n/2 - i)
		result[n-1-i] = (n/2 - i)
	}

	return result
}