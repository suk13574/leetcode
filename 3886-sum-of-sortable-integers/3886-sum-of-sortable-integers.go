func sortableIntegers(nums []int) int {
	n := len(nums)

	divisors := []int{}

	for k := 1; k*k <= n; k++ {
		if n%k == 0 {
			divisors = append(divisors, k)

			if k != n/k {
				divisors = append(divisors, n/k)
			}
		}
	}

	res := 0

	for _, k := range divisors {
		prevMax := 0
		sortable := true

		for l := 0; l < n; l += k {
			r := l + k
			sub := nums[l:r]

			cnt := 0
			minVal, maxVal := sub[0], sub[0]

			for i := 0; i < k; i++ {
				if sub[i] > sub[(i+1)%k] {
					cnt++
				}

				if cnt >= 2 {
					sortable = false
					break
				}

				minVal = min(minVal, sub[i])
				maxVal = max(maxVal, sub[i])
			}

			if !sortable {
				break
			}

			if prevMax > minVal {
				sortable = false
				break
			}

			prevMax = maxVal
		}

		if sortable {
			res += k
		}
	}

	return res
}