func arrayRankTransform(arr []int) []int {
	n := len(arr)

	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	rankMap := make(map[int]int)
	rank := 1

	for _, v := range sorted {
		if _, ok := rankMap[v]; !ok {
			rankMap[v] = rank
			rank++
		}
	}

	res := make([]int, n)
	for i, v := range arr {
		res[i] = rankMap[v]
	}

	return res
}