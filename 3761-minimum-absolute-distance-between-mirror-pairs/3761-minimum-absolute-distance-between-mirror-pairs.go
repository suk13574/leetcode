import "math"

func minMirrorPairDistance(nums []int) int {
	lastSeen := make(map[int]int)

	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		idx, ok := lastSeen[nums[i]]
		if ok {
			res = min(res, i-idx)
		}

		lastSeen[reverse(nums[i])] = i
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func reverse(x int) int {
	res := 0
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}

	return res
}