func minOperations(nums []int) int {
	stack := []int{}

	pop := func() {
		n := len(stack)
		stack = stack[:n-1]
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		n := len(stack)

		for n != 0 && stack[n-1] > nums[i] {
			pop()
			n = len(stack)
			res++
		}

		if n != 0 && stack[n-1] == nums[i] {
			continue
		}
		if nums[i] == 0 {
			continue
		}

		stack = append(stack, nums[i])
	}

	res += len(stack)

	return res
}