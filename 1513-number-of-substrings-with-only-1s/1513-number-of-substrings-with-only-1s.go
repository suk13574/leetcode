func numSub(s string) int {
	const MOD = 1_000_000_007

	res := 0
	ones := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
            res += (ones % MOD)
			res %= MOD
		}
		if s[i] == '0' {
			ones = 0
		}
	}

	return res % MOD
}