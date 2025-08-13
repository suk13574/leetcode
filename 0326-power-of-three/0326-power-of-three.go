import "math"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	var max3Pow int64 = 1

	for max3Pow <= math.MaxInt64/3 {
		max3Pow *= 3
	}

	return max3Pow%int64(n) == 0

}