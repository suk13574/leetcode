func countValidSelections(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	sumLeft := 0
	res := 0

	for i := 0; i < len(nums); i++ {
		sumRight := total - sumLeft

		if nums[i] == 0 {
			switch diff := sumLeft - sumRight; diff {
			case 0:
				res += 2
			case 1, -1:
				res += 1

			}
		}

		sumLeft += nums[i]
	}

	return res
}
