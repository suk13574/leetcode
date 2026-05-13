func minMoves(nums []int, limit int) int {
	n := len(nums)

	minT := 2
	maxT := limit * 2

	diff := make([]int, maxT+2)
	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		if a > b {
			a, b = b, a
		}

		diff[a+1]--
		diff[a+b]--
		diff[a+b+1]++
		diff[b+limit+1]++
	}

	cur := n
	res := n
	for T := minT; T <= maxT; T++ {
		cur += diff[T]
		res = min(res, cur)
	}
	
	return res
}