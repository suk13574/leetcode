func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := len(nums)-1; i >= 2; i-- {
		a, b, c := nums[i-2], nums[i-1], nums[i]
		if a+b > c {
			return a + b + c
		}

	}

	return 0
}