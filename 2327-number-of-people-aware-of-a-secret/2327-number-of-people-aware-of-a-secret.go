func peopleAwareOfSecret(n int, delay int, forget int) int {
	const MOD = 1000000007

	newKnown := make([]int, n+1)
	newKnown[1] = 1

	share := 0

	for i := 2; i <= n; i++ {
		if i-delay >= 1 {
			share += newKnown[i-delay]
			if share > MOD {
				share -= MOD
			}
		}

		if i-forget >= 1 {
			share -= newKnown[i-forget]
			if share < 0 {
				share += MOD
			}
		}

		newKnown[i] = share
	}

	result := 0
	start := n - forget + 1
	if start < 1 {
		start = 1
	}

	for i := start; i <= n; i++ {
		result += newKnown[i]
		if result > MOD {
			result -= MOD
		}
	}

	return result
}