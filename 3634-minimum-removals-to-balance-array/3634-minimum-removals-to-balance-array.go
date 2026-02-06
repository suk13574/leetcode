func minRemoval(nums []int, k int) int {
	sort.Ints(nums)
	i := 0
	maxLen := 0

	for j := 0; j < len(nums); j++ {
		for nums[i]*k < nums[j] {
			i++
		}

		maxLen = max(maxLen, j-i+1)
	}

	return len(nums) - maxLen
}