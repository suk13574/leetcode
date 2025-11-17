func kLengthApart(nums []int, k int) bool {
	lastOne := -1
	for i, n := range nums {
		if n == 1 {
			if lastOne != -1 && i-lastOne-1 < k {
				return false
			}
			lastOne = i
		}
	}
	return true
}