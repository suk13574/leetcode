func uniqueXorTriplets(nums []int) int {
	const limit = 1 << 11

	seen := make([]bool, limit)
	unique := make([]int, 0, len(nums))

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	pairSeen := make([]bool, limit)
	pairs := make([]int, 0, limit)
	for i := 0; i < len(unique); i++ {
		for j := i; j < len(unique); j++ {
			x := unique[i] ^ unique[j]
			if !pairSeen[x] {
				pairSeen[x] = true
				pairs = append(pairs, x)
			}
		}
	}

	resSeen := make([]bool, limit)
	res := 0
	for _, num := range unique {
		for _, x := range pairs {
			triple := num ^ x
			if !resSeen[triple] {
				resSeen[triple] = true
				res++
			}
		}
	}

	return res
}