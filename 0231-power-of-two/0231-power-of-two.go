import "math"

func isPowerOfTwo(n int) bool {
	pow := make(map[float64]int)
	for i := 0; i < 31; i++ {
		pow[math.Pow(2, float64(i))] = i
	}

	if _, ok := pow[float64(n)]; ok {
		return true
	}
	return false
}