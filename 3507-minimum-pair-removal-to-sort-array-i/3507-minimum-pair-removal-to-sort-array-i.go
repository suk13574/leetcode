func minimumPairRemoval(nums []int) int {
	isSort := func(n []int) bool {
		for i := 1; i < len(n); i++ {
			if n[i-1] > n[i] {
				return false
			}
		}
		return true
	}

	res := 0
	for !isSort(nums) {
		res++

		idx := 1
		sum := nums[0] + nums[1]
		for i := 2; i < len(nums); i++ {
			s := nums[i] + nums[i-1]
			if sum > s {
				sum = s
				idx = i
			}
		}

		nums[idx-1] = sum
		nums = append(nums[:idx], nums[idx+1:]...)
	}
	return res
}