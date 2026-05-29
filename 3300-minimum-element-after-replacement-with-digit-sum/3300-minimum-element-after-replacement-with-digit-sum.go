func minElement(nums []int) int {
	sumDigits := func(x int) int {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}

		return sum
	}

	res := 1 << 30
	for _, num := range nums {
		res = min(res, sumDigits(num))
	}

	return res
}