func numOfSubarrays(arr []int) int {
	const MOD int64 = 1_000_000_007

	odd := int64(0)
	even := int64(1)

	sum := int64(0)
	for _, v := range arr {
		sum += int64(v)

		if sum&1 == 1 {
			odd++
		} else {
			even++
		}
	}

	return int((odd * even) % MOD)
}