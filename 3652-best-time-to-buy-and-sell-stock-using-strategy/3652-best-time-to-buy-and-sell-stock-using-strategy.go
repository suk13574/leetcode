func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	prefixSum := make([]int64, n+1)

	var totalProfit int64 = 0
	for i := 0; i < n; i++ {
		profit := prices[i] * strategy[i]
		totalProfit += int64(profit)
		prefixSum[i+1] = totalProfit
	}

	var kRange int64 = 0
	var res int64 = totalProfit
	l := 0
	for r := k / 2; r < n; r++ {
		kRange += int64(prices[r])
		if (r - l + 1) < k {
			continue
		}

		base := totalProfit - prefixSum[r+1] + prefixSum[l]
		profit := base + kRange
		res = max(res, profit)

		kRange -= int64(prices[(r+l)/2])
		l++
	}

	return res
}