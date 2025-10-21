import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	Ls := make([]int, n)
	Rs := make([]int, n)
	for i := 0; i < n; i++ {
		Ls[i] = nums[i] - k
		Rs[i] = nums[i] + k
	}

	i, j := 0, 0
	cur, maxCover := 0, 0
	for i < n && j < n {
		if Ls[i] <= Rs[j] {
			cur++
			maxCover = max(maxCover, cur)
			i++
		} else {
			cur--
			j++
		}
	}
	res := min(maxCover, numOperations)

	L, R := 0, -1
	for gStart := 0; gStart < n; {
		gEnd := gStart
		for gEnd+1 < n && nums[gEnd+1] == nums[gStart] {
			gEnd++
		}
		v := nums[gStart]
		lo, hi := v-k, v+k

		for R+1 < n && nums[R+1] <= hi {
			R++
		}
		for L <= R && nums[L] < lo {
			L++
		}

		coverAtV := R - L + 1
		cV := gEnd - gStart + 1
		cand := min(coverAtV, cV+numOperations)
		res = max(res, cand)

		gStart = gEnd + 1
	}

	return res
}