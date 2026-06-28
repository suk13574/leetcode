func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)

	cur := 0

	for _, v := range arr {
		if v > cur {
			cur++
		}
	}

	return cur
}