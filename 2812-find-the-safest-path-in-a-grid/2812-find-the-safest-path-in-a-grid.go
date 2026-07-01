import "container/heap"

type Node struct {
	x int
	y int
	d int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].d > pq[j].d }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return node
}

func maximumSafenessFactor(grid [][]int) int {
	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	q := []Node{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, Node{i, j, 0})
				visited[i][j] = true
				dist[i][j] = 0
			}
		}
	}

	for cur := 0; cur < len(q); cur++ {
		node := q[cur]

		for _, move := range moves {
			nx := node.x + move[0]
			ny := node.y + move[1]

			if nx < 0 || ny < 0 || nx >= m || ny >= n || visited[nx][ny] {
				continue
			}

			nd := node.d + 1
			dist[nx][ny] = nd
			q = append(q, Node{nx, ny, nd})
			visited[nx][ny] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{0, 0, dist[0][0]})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if visited[node.x][node.y] {
			continue
		}

		visited[node.x][node.y] = true

		if node.x == m-1 && node.y == n-1 {
			return node.d
		}

		for _, move := range moves {
			nx := node.x + move[0]
			ny := node.y + move[1]

			if nx < 0 || ny < 0 || nx >= m || ny >= n || visited[nx][ny] {
				continue
			}

			nd := min(node.d, dist[nx][ny])
			heap.Push(&pq, &Node{nx, ny, nd})
		}

	}

	return 0
}