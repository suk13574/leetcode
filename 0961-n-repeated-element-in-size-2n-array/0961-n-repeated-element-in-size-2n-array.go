func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	seen := make(map[int]bool, n+1)
	for _, n := range nums {
		if seen[n] {
			return n
		}

		seen[n] = true
	}

	return 0
}