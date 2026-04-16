func solveQueries(nums []int, queries []int) []int {
	idxs := make(map[int][]int)
	for i, num := range nums {
		if _, ok := idxs[num]; !ok {
			idxs[num] = []int{}
		}
		idxs[num] = append(idxs[num], i)
	}

	binarySearch := func(q int, pos []int) int {
		l := 0
		r := len(pos)

		for l < r {
			mid := (l + r) / 2
			if pos[mid] < q {
				l = mid + 1
			} else {
				r = mid
			}
		}

		return l
	}

	n := len(nums)
	m := len(queries)
	res := make([]int, m)

	for j := 0; j < m; j++ {
		q := queries[j]
		positions := idxs[nums[q]]
		l := len(positions)

		if l == 1 {
			res[j] = -1
			continue
		}

		i := binarySearch(q, positions)
		right := positions[(i+1)%l]
		d1 := abs(q - right)
		d1 = min(d1, n-d1)

		left := positions[(i-1+l)%l]
		d2 := abs(q - left)
		d2 = min(d2, n-d2)

		res[j] = min(d1, d2)
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}