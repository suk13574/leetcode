type Node struct {
	x, y int
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0

	bfs := func(i, j int) {
		isClosed := true

		queue := []Node{{i, j}}
		visited[i][j] = true

		for pos := 0; pos < len(queue); pos++ {
			node := queue[pos]
			x, y := node.x, node.y

			if x == 0 || y == 0 || x == m-1 || y == n-1 {
				isClosed = false
			}

			for _, move := range moves {
				nx := x + move[0]
				ny := y + move[1]

				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					continue
				}

				if grid[nx][ny] == 1 || visited[nx][ny] {
					continue
				}

				queue = append(queue, Node{nx, ny})
				visited[nx][ny] = true
			}
		}

		if isClosed {
			res++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				bfs(i, j)
			}
		}
	}

	return res
}