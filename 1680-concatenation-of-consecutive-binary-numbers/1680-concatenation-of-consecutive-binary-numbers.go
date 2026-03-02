func concatenatedBinary(n int) int {
	const MOD = 1000000007

	lenBinary := func(num int) int {
		if num == 0 {
			return 1
		}
		l := 0
		for num > 0 {
			l++
			num >>= 1
		}
		return l
	}

	res := 0
	for i := 1; i <= n; i++ {
		l := lenBinary(i)
		res <<= l
		res |= i
		res %= MOD
	}

	return res
}