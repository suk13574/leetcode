func numOfSubarrays(arr []int) int {
	const MOD = 1_000_000_007

	sum := 0
	odd := 0
	even := 1
	res := 0

	for _, v := range arr {
		sum += v

		if sum&1 == 1 {
			res = (res + even) % MOD
            odd++
		} else {
			res = (res + odd) % MOD
            even++
		}
	}

	return res % MOD
}