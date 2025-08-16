func maximum69Number(n int) int {
	var result int
	count := 1

	for i := 1000; i > 0; i = i / 10 {
		q := n / i
		if q == 0 {
			continue
		}
		if q == 6 && count == 1 {
			count--
			q = 9
		}
		result += q * i
		n = n % i
	}

	return result
}