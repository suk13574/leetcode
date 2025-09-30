func triangularSum(nums []int) int {
	const MOD = 10

	n := len(nums)
	count := make([]int, n)
	count[0] = 1

	for length := 1; length < n; length++ {
		for i := length; i > 0; i-- {
			count[i] = (count[i] + count[i-1]) % MOD
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += (nums[i] * count[i]) % MOD
	}
	return res % MOD
}