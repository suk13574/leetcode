func findSmallestInteger(nums []int, value int) int {
	mod := make([]int, value)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		r := num % value
		if r < 0 {
			r += value
		}
		mod[r]++
	}

	mex := 0
	for {
		r := mex % value
		if mod[r] == 0 {
			return mex
		}
		mod[r]--
		mex++
	}
}