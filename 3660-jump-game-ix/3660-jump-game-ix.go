func maxValue(nums []int) []int {
	n := len(nums)

	prefixMax := make([]int, n) // max of nums[0..i]
	suffixMin := make([]int, n) // min of nums[i..n-1]

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}

	res := make([]int, n)
	l := 0
	best := 0
	for r := 0; l <= r && r < n; r++ {
		best = max(best, prefixMax[r])

		if r == n-1 || prefixMax[r] <= suffixMin[r+1] {
			for l <= r {
				res[l] = best
				l++
			}
			best = 0
		}
	}

	return res
}