func remainingMethods(n int, k int, invocations [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, invocation := range invocations {
		u, v := invocation[0], invocation[1]

		graph[u] = append(graph[u], v)
	}

	suspicious := make([]bool, n)
	visited := make([]bool, n)
	queue := []int{k}
	visited[k] = true

	for pos := 0; pos < len(queue); pos++ {
		node := queue[pos]

		suspicious[node] = true

		for _, v := range graph[node] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	for _, invocation := range invocations {
		u, v := invocation[0], invocation[1]

		if !suspicious[u] && suspicious[v] {
			res := make([]int, n)
			for i := 0; i < n; i++ {
				res[i] = i
			}
			return res
		}
	}

	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !suspicious[i] {
			res = append(res, i)
		}
	}

	return res
}