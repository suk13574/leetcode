func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)

	prevEnd := -1
	count := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			if count >= 2*k {
				return true
			}
			if count >= k {
				prevEnd = i - 1
			}
            count = 1
		}

		if count == k {
			start := i - k + 1
			if prevEnd >= 0 && prevEnd+1 == start {
				return true
			}
		}
		if count >= 2*k {
			return true
		}
	}
	return count >= 2*k
}