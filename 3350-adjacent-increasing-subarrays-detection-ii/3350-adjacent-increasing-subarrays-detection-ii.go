func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	increLen := make([]int, n)
	increLen[0] = 1

	res := 0
	prevLen := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			increLen[i] = increLen[i-1] + 1
		} else {
			increLen[i] = 1
			prevLen = increLen[i-1]
		}

		k := max(max(prevLen, increLen[i])/2, min(prevLen, increLen[i]))
		res = max(k, res)
	}
	return res
}