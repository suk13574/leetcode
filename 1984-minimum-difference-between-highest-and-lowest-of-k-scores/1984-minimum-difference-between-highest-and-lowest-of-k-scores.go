func minimumDifference(nums []int, k int) int {
	n := len(nums)
	if k == 1 {
		return 0
	}

	sort.Ints(nums)
	res := 1 << 30
	for i := 0; i+k-1 < n; i++ {
		diff := nums[i+k-1] - nums[i]
		res = min(res, diff)
	}

	return res

}