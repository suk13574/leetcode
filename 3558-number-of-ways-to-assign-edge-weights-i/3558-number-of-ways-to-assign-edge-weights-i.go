const MOD int = 1_000_000_007

type Node struct {
	val   int
	depth int
}

func assignEdgeWeights(edges [][]int) int {
	n := len(edges) + 1
	graph := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = []int{}
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	queue := []Node{{1, 0}}

	visited := make([]bool, n+1)
	visited[1] = true

	maxDepth := 0
	for cur := 0; cur < len(queue); cur++ {
		node := queue[cur]
		maxDepth = max(maxDepth, node.depth)

		for _, next := range graph[node.val] {
			if visited[next] {
				continue
			}

			visited[next] = true
			queue = append(queue, Node{next, node.depth + 1})
		}
	}

	return modPow(2, maxDepth-1)
}

func modPow(base, exp int) int {
	res := 1

	for exp > 0 {
		if exp&1 == 1 {
			res = res * base % MOD
		}

		base = base * base % MOD
		exp >>= 1
	}

	return res
}