func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007
	n := len(s)

	prefixSum := make([]int, n+1)
	prefixCnt := make([]int, n+1)
	digits := make([]int, 0, n)
	for i := 0; i < n; i++ {
		d := int(s[i] - '0')

		prefixSum[i+1] = prefixSum[i] + d
		prefixCnt[i+1] = prefixCnt[i]

		if d != 0 {
			prefixCnt[i+1]++
			digits = append(digits, d)
		}
	}

	k := len(digits)

	pow10 := make([]int64, k+1)
	pow10[0] = 1
	for i := 1; i <= k; i++ {
		pow10[i] = pow10[i-1] * 10 % MOD
	}

	prefVal := make([]int64, k+1)
	for i := 0; i < k; i++ {
		prefVal[i+1] = (prefVal[i]*10 + int64(digits[i])) % MOD
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]

		sum := prefixSum[r+1] - prefixSum[l]

		left := prefixCnt[l]
		right := prefixCnt[r+1]

		if left == right {
			res[i] = 0
			continue
		}

		length := right - left
		x := (prefVal[right] - prefVal[left]*pow10[length]) % MOD
		if x < 0 {
			x += MOD
		}
		res[i] = int(x * int64(sum) % MOD)
	}

	return res
}