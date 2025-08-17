import "math"

func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1

	var prevProb float64

	if k > 0 {
		prevProb = 1
	} else {
		prevProb = 0
	}

	for i := 1; i <= n; i++ {
		dp[i] = prevProb / float64(maxPts)

		if i < k {
			prevProb += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			prevProb -= dp[i-maxPts]
		}
	}

	var result float64
	for _, v := range dp[k:] {
		result += v
	}

	return math.Round(result*1e5) / 1e5
}