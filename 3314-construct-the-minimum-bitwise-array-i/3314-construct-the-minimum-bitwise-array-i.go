func minBitwiseArray(nums []int) []int {
	n := len(nums)

	find := func(v int) int {
		for i := 0; i < v; i++ {
			if i|(i+1) == v {
				return i
			}
		}
		return -1
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = find(nums[i])

	}

	return res
}