func makeTheIntegerZero(num1 int, num2 int) int {
	countSum := func(num int) int {
		result := 0
		for num > 0 {
			if num&1 == 1 {
				result++
			}
			num = num >> 1
		}

		return result
	}

	var leftTerm int
	for i := 1; i < 61; i++ {
		leftTerm = num1 - (i * num2)

		if leftTerm >= i && countSum(leftTerm) <= i {
			return i
		}
	}

	return -1
}