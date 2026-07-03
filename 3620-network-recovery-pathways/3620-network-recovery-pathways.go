type Node struct {
	v    int
	cost int
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)

	graph := make(map[int][]Node)
	indeg := make([]int, n)
	maxCost := 0
	for _, edge := range edges {
		u, v, c := edge[0], edge[1], edge[2]

		if !online[u] || !online[v] {
			continue
		}

		if _, ok := graph[u]; !ok {
			graph[u] = []Node{}
		}

		graph[u] = append(graph[u], Node{v, c})
		indeg[v]++
		maxCost = max(maxCost, c)
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := make([]int, 0, n)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		topo = append(topo, u)

		for _, next := range graph[u] {
			v := next.v
			indeg[v]--

			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	can := func(minimum int) bool {
		const INF int64 = 1 << 60

		dist := make([]int64, n)
		for i := 0; i < n; i++ {
			dist[i] = INF
		}

		dist[0] = 0

		for _, u := range topo {
			if dist[u] == INF {
				continue
			}

			if !online[u] {
				continue
			}

			for _, next := range graph[u] {
				v := next.v
				cost := next.cost

				if cost < minimum {
					continue
				}

				if !online[v] {
					continue
				}

				newCost := dist[u] + int64(cost)
				dist[v] = min(dist[v], newCost)
			}
		}

		return dist[n-1] <= k
	}

	if !can(0) {
		return -1
	}

	l, r := 0, maxCost
	res := 0

	for l <= r {
		mid := l + (r-l)/2

		if can(mid) {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}