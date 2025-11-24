func prefixesDivBy5(nums []int) []bool {
	res := make([]bool, len(nums))
	mod := 0

	for i := 0; i < len(nums); i++ {
		mod = (mod*2 + nums[i]) % 5
		res[i] = mod == 0
	}

	return res
}