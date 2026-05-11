import "slices"

func separateDigits(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		x := num

		digits := []int{}
		for x > 0 {
			digits = append(digits, x%10)
			x /= 10
		}

		slices.Reverse(digits)
		res = append(res, digits...)
	}

	return res
}