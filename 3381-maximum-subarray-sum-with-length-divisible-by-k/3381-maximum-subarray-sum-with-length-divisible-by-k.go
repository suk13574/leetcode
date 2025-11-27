func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := make([]int64, n+1)

	// prefixSum[i] = nums[0] + ... + nums[i-1]
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + int64(nums[i])
	}

	const INF = math.MaxInt64

	minPref := make([]int64, k)
	for i := 0; i < k; i++ {
		minPref[i] = INF
	}
	minPref[0] = 0 // prefix[0] = 0, 0%k = 0

	res := int64(math.MinInt64)
	for r := 1; r <= n; r++ {
		mod := r % k

		if minPref[mod] != INF {
			sumSub := prefixSum[r] - minPref[mod]
			res = max(res, sumSub)
		}

		if prefixSum[r] < minPref[mod] {
			minPref[mod] = prefixSum[r]
		}
	}
	return res
}
