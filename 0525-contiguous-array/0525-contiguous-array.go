func findMaxLength(nums []int) int {
	first := make(map[int]int)
	first[0] = -1

	sum := 0
	res := 0
	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}

		if idx, ok := first[sum]; ok {
			res = max(res, i-idx)
		} else {
			first[sum] = i
		}
	}

	return res
}