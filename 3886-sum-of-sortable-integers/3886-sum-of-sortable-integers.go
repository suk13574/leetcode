func sortableIntegers(nums []int) int {
	n := len(nums)
	res := 0

	for k := 1; k*k <= n; k++ {
		if n%k != 0 {
			continue
		}

		if isSortable(nums, k) {
			res += k
		}

		if k != n/k && isSortable(nums, n/k) {
			res += n / k
		}
	}

	return res
}

func isSortable(nums []int, k int) bool {
	prevMax := 0

	for l := 0; l < len(nums); l += k {
		sub := nums[l : l+k]
		minVal, maxVal := sub[0], sub[0]
		descents := 0

		for i, v := range sub {
			if sub[i] > sub[(i+1)%k] {
				descents++

				if descents > 1 {
					return false
				}
			}

            minVal = min(minVal, v)
            maxVal = max(maxVal, v)
		}

        if prevMax > minVal {
            return false
        }

        prevMax = maxVal        
	}

	return true
}