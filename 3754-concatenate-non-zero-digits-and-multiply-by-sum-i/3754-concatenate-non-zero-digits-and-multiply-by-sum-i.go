func sumAndMultiply(n int) int64 {
	x := 0
	sum := 0

	for n > 0 {
		tail := n % 10
		n /= 10

		if tail == 0 {
			continue
		}

		x = x*10 + tail
		sum += tail
	}

	tmp := 0
	for x > 0 {
		tmp = tmp*10 + (x % 10)
		x /= 10
	}
	x = tmp

	return int64(x) * int64(sum)
}