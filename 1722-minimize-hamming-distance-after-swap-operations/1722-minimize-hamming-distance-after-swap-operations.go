func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, swap := range allowedSwaps {
		a, b := swap[0], swap[1]
		if a > b {
			a, b = b, a
		}

		union(parent, a, b)
	}

	groups := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		p := find(parent, i)

		if _, ok := groups[p]; !ok {
			groups[p] = make(map[int]int)
		}
		groups[p][source[i]]++
	}

	res := 0
	for i := 0; i < n; i++ {
		p := find(parent, i)
		if c, ok := groups[p][target[i]]; !ok || c <= 0 {
			res++
		}
		groups[p][target[i]]--
	}
	return res
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}

	return parent[x]
}

func union(parent []int, x, y int) {
	xp := find(parent, x)
	yp := find(parent, y)

	parent[yp] = xp
}