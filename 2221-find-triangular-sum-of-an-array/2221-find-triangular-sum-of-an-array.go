func triangularSum(nums []int) int {
	const MOD = 10

	n := len(nums)
	for length := n; length > 1; length-- {
		for i := 0; i < length-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % MOD
		}
	}
	return nums[0]
}