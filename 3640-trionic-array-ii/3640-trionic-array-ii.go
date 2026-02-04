func maxSumTrionic(nums []int) int64 {
	const NEG int64 = -(1 << 62)
	n := len(nums)

	dp1 := make([]int64, n)
	dp2 := make([]int64, n)
	dp3 := make([]int64, n)
	for i := 0; i < n; i++ {
		dp1[i], dp2[i], dp3[i] = NEG, NEG, NEG
	}

	res := NEG
	for i := 1; i < n; i++ {
		a := int64(nums[i-1])
		b := int64(nums[i])
		if nums[i-1] < nums[i] {
			dp1[i] = max(dp1[i], a+b)
			if dp1[i-1] != NEG {
				dp1[i] = max(dp1[i], dp1[i-1]+b)
			}

			if dp3[i-1] != NEG {
				dp3[i] = max(dp3[i], dp3[i-1]+b)
			}
			if dp2[i-1] != NEG {
				dp3[i] = max(dp3[i], dp2[i-1]+b)
			}
		}

		if nums[i-1] > nums[i] {
			if dp2[i-1] != NEG {
				dp2[i] = max(dp2[i], dp2[i-1]+b)
			}
			if dp1[i-1] != NEG {
				dp2[i] = max(dp2[i], dp1[i-1]+b)
			}
		}

		res = max(res, dp3[i])
	}

	return res
}