func minBitwiseArray(nums []int) []int {
	makeBinary := func(v int) []int {
		binary := []int{}
		for v > 0 {
			binary = append(binary, v%2)
			v = v / 2
		}
		binary = append(binary, 0)
		slices.Reverse(binary)
		return binary
	}

	makeDecimal := func(v []int) int {
		d := 0
		base := 1

		for i := len(v) - 1; i >= 0; i-- {
			d += v[i] * base
			base *= 2
		}

		return d
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 {
			res[i] = -1
            continue
		}

		binary := makeBinary(nums[i])
		for i := len(binary) - 1; i >= 1; i-- {
			if binary[i-1] == 0 {
				binary[i] = 0
				break
			}
		}

		res[i] = makeDecimal(binary)

	}
	return res
}