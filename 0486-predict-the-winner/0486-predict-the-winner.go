func predictTheWinner(nums []int) bool {
	n := len(nums)
	if n&1 == 0 {
		return true
	}

	var maxDiff func(i, j int) int
	maxDiff = func(i, j int) int {
        if i == j {
            return nums[i]
        }

		return max(
			nums[i]-maxDiff(i+1, j),
			nums[j]-maxDiff(i, j-1),
		)
	}

	return maxDiff(0, n-1) >= 0
}