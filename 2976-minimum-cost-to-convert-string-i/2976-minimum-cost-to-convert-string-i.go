func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const INF = 1 << 60
	graph := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		graph[i] = make([]int64, 26)
		for j := 0; j < 26; j++ {
			graph[i][j] = INF
		}
		graph[i][i] = 0

	}

	for i := 0; i < len(original); i++ {
		src := original[i] - 'a'
		dst := changed[i] - 'a'
		graph[src][dst] = min(graph[src][dst], int64(cost[i]))
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	var res int64
	for i := 0; i < len(source); i++ {
		s := source[i] - 'a'
		d := target[i] - 'a'

		if graph[s][d] == INF {
			return -1
		}
		res += graph[s][d]
	}

	return res
}