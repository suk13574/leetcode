func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)

	lcms := make([]int64, 1<<n)
	lcms[0] = 1

	for mask := 1; mask < (1 << n); mask++ {
		bit := mask & -mask
		idx := bits.TrailingZeros(uint(bit))
		prev := mask ^ bit

		lcms[mask] = lcm(lcms[prev], int64(coins[idx]))
	}

	count := func(x int64) int64 {
		var res int64

		for mask := 1; mask < (1 << n); mask++ {
			n := bits.OnesCount(uint(mask))

			if n&1 == 1 {
				res += x / lcms[mask]
			} else {
				res -= x / lcms[mask]
			}
		}

		return res
	}

	minCoin := coins[0]
	for i := 1; i < n; i++ {
		minCoin = min(minCoin, coins[i])
	}

	l := int64(1)
	r := int64(minCoin) * int64(k)

	for l < r {
		mid := l + (r-l)/2

		if count(mid) >= int64(k) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l

}