func maxFrequencyElements(nums []int) int {
	maxCount := 0
	numCount := make(map[int]int)

	for _, n := range nums {
		numCount[n]++
		maxCount = max(maxCount, numCount[n])
	}

	res := 0
	for _, c := range numCount {
		if c == maxCount {
			res += c
		}
	}

	return res
}