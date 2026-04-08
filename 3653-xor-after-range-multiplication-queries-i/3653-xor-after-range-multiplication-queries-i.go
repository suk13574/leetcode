func xorAfterQueries(nums []int, queries [][]int) int {
	const MOD = 1000000007
	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]

		for idx := l; idx <= r; idx = idx + k {
			nums[idx] = (nums[idx] * v) % MOD
		}
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}