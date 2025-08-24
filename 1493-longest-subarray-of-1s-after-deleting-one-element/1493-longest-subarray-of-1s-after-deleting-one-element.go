func longestSubarray(nums []int) int {
	result := 0

	left := 0
	zeroIdx := -1

	count := 0
	for r, num := range nums {
		if num == 0 {
			if zeroIdx == -1 {
				zeroIdx = r
			} else {
				result = max(result, count)
				count -= (zeroIdx - left)

				left = zeroIdx + 1
				zeroIdx = r
			}
		} else {
			count++
		}
	}

	if zeroIdx == -1 {
		return count - 1
	}

	return max(result, count)
}
