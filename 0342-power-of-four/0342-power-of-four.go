import (
	"math"
)

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 1; i < math.MaxInt64/4; i = i * 4 {
		if i == n {
			return true
		}
	}

	return false
}