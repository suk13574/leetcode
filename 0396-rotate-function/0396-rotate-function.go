func maxRotateFunction(nums []int) int {
	n := len(nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	prev := 0
	for i := 0; i < n; i++ {
		prev += i * nums[i]
	}

	res := prev
	for i := 1; i < n; i++ {
		rotate := prev + total - n*nums[n-i]
		res = max(res, rotate)
		prev = rotate
	}
	
	return res
}