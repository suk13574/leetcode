import "container/heap"

type Node struct {
	x      int
	y      int
	health int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].health > pq[j].health }

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

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Node{0, 0, health - grid[0][0]})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		if visited[node.x][node.y] {
			continue
		}
		visited[node.x][node.y] = true

		if node.health <= 0 {
			return false
		}

		if node.x == m-1 && node.y == n-1 {
			return true
		}

		for _, move := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nx := node.x + move[0]
			ny := node.y + move[1]
			
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			nHealth := node.health - grid[nx][ny]

			heap.Push(&pq, &Node{nx, ny, nHealth})
		}
	}

	return false
}