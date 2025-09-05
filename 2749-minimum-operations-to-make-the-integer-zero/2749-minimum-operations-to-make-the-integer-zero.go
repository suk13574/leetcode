func makeTheIntegerZero(num1 int, num2 int) int {
	countBit := func(num int) int {
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
	for i := 1; ; i++ {
		leftTerm = num1 - (i * num2)

		if leftTerm < i {
			return -1
		}

		if leftTerm >= i && countBit(leftTerm) <= i {
			return i
		}
	}
}