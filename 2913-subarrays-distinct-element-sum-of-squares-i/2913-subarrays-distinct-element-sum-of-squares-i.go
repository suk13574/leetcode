func sumCounts(nums []int) int {
	n := len(nums)

	res := 0

	for l := 0; l < n; l++ {
		hash := make(map[int]bool)
		cnt := 0

		for r := l; r < n; r++ {
			if !hash[nums[r]] {
				hash[nums[r]] = true
				cnt++
			}

			res += cnt * cnt
		}
	}

	return res
}