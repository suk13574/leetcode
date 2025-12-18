func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	prefixSum := make([]int64, n+1)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + int64(strategy[i])*int64(prices[i])

	}
	var totalProfit = prefixSum[n]

	var kHalfSum int64 = 0
	for i := k / 2; i < k; i++ {
		kHalfSum += int64(prices[i])
	}

	var res int64 = prefixSum[n]
	for l := 0; l+k <= n; l++ {
		r := l + k - 1

		kWindow := prefixSum[r+1] - prefixSum[l]
		profit := totalProfit - kWindow + kHalfSum
		res = max(res, profit)

		if l+k < n {
			kHalfSum -= int64(prices[l+k/2])
			kHalfSum += int64(prices[l+k])

		}
	}

	return res
}