func rotatedDigits(n int) int {
	isGoodNumber := func(x int) bool {
		res := false

		for x > 0 {
			target := x % 10

			switch target {
			case 3, 4, 7:
				return false
			case 2, 5, 6, 9:
				res = true
			}

			x /= 10
		}

		return res
	}

	res := 0
	for i := 1; i <= n; i++ {
		if isGoodNumber(i) {
			res++
		}
	}

	return res
}