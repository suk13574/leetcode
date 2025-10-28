func countValidSelections(nums []int) int {
	n := len(nums)

	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}

	left := 0
	right := total
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			left += nums[i]
			right -= nums[i]
			continue
		}

		if left == right {
			res += 2
		} else if left+1 == right || right+1 == left {
			res += 1
		}

	}
	return res
}