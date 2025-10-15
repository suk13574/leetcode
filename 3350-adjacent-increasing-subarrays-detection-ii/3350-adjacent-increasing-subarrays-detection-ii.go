func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)

	res := 0
	nowLen := 1
	prevLen := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			nowLen++
		} else {
			res = max(res, nowLen/2)
			prevLen = nowLen
			nowLen = 1
		}

		res = max(res, min(prevLen, nowLen))
	}

	res = max(res, nowLen/2)

	return res
}