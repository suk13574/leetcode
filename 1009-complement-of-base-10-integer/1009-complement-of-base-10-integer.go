func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	base := 1

	for base < n {
		base <<= 1
	}

	if base == n {
		return base - 1
	} else {
		return base - 1 - n
	}
}