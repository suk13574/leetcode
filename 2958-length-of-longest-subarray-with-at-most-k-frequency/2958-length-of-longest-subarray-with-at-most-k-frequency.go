func maxSubarrayLength(nums []int, k int) int {
	count := make(map[int]int)

	res := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		v := nums[r]

		count[v]++

		for count[v] > k {
			count[nums[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}