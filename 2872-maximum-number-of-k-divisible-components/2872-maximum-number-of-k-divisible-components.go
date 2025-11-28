func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	tree := make([][]int, n)
	for i := 0; i < n; i++ {
		tree[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}

	res := 0

	var dfs func(node, parent int) int
	dfs = func(u, parent int) int {
		sum := values[u]

		for _, v := range tree[u] {
			if v == parent {
				continue
			}
			childSum := dfs(v, u)
			sum += childSum
		}

		if sum%k == 0 {
			res++
			return 0
		} else {
			return sum
		}
	}

	dfs(0, -1)
	return res
}