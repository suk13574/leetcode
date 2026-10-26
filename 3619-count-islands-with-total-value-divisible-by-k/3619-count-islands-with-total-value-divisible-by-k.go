type Node struct {
	x, y int
}

func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visited := make([][]bool, m)

	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0

	bfs := func(i, j int) {
		queue := []Node{{i, j}}
		visited[i][j] = true

		sum := grid[i][j]

		for pos := 0; pos < len(queue); pos++ {
			node := queue[pos]

			for _, move := range moves {
				nx := node.x + move[0]
				ny := node.y + move[1]

				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}

				if grid[nx][ny] == 0 || visited[nx][ny] {
					continue
				}

				queue = append(queue, Node{nx, ny})
				visited[nx][ny] = true
				sum += grid[nx][ny]
			}
		}

		if sum%k == 0 {
			res ++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 && !visited[i][j] {
				bfs(i, j)
			}
		}
	}

	return res
}