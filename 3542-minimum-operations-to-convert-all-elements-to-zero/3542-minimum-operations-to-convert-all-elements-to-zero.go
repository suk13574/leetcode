func minOperations(nums []int) int {
	stack := []int{}

	res := 0
	for _, x := range nums {
		if x == 0 {
			res += len(stack)
			stack = stack[:0]
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > x {
			stack = stack[:len(stack)-1]
			res++
		}

		if len(stack) == 0 || stack[len(stack)-1] < x {
			stack = append(stack, x)
		}
	}

	res += len(stack)
	return res
}