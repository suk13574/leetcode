var dp = map[int]int{0: 0}

func OneBitOperation(n int) int {
	if val, ok := dp[n]; ok {
		return val
	}

	i := 0 // index of the leftmost 1-bit
	for (1 << (i + 1)) <= n {
		i++
	}
	msb := 1 << i // value with only the leftmost 1-bit set (2^i)
	r := n - msb

	res := (1<<(i+1) - 1) - OneBitOperation(r)
	dp[n] = res
	return res
}

func minimumOneBitOperations(n int) int {
	return OneBitOperation(n)
}