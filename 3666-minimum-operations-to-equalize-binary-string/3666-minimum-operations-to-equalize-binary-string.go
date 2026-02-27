func minOperations(s string, k int) int {
	n := len(s)
	zeroCnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			zeroCnt++
		}
	}

	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = -1
	}

	parent := make([]int, n+4)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if x > len(parent) {
			return x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	erase := func(x int) {
		parent[x] = find(x + 2)
	}

	queue := make([]int, 0)
	queue = append(queue, zeroCnt)
	dist[zeroCnt] = 0
	erase(zeroCnt)

	for head := 0; head < len(queue); head++ {
		z := queue[head]

		if z == 0 {
			return dist[z]
		}

		lo := max(0, k-(n-z))
		hi := min(k, z)

		// nextZeroCnt = z + k - 2*i
		L := z + k - (2 * hi)
		R := z + k - (2 * lo)

		x := find(L)
		for x <= R && x <= n {
			dist[x] = dist[z] + 1
			queue = append(queue, x)

			erase(x)
			x = find(x)
		}
	}

	return -1
}