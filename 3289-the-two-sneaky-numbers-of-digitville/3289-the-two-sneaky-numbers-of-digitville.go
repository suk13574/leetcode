func getSneakyNumbers(nums []int) []int {
	count := make([]int, 101)
	res := []int{}

	for _, n := range nums {
		if count[n] == 1 {
			res = append(res, n)
		}
		count[n]++
	}
	return res
}