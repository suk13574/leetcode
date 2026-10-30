import (
	"container/heap"
	"math"
)

type Edge struct {
	v, w int
}

type State struct {
	node  int
	cost  int
	count int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	state := x.(*State)
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]

	return state
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	graph := make([][]Edge, n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, k+1)

		for j := 0; j <= k; j++ {
			dist[i][j] = math.MaxInt
		}
	}

	dist[0][1] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &State{node: 0, cost: 0, count: 1}

	heap.Init(&pq)

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*State)

		if cur.cost > dist[cur.node][cur.count] {
			continue
		}

		if cur.node == n-1 {
			return cur.cost
		}

		for _, e := range graph[cur.node] {
			next := e.v
			nextCost := cur.cost + e.w
			nextCount := cur.count

			if labels[cur.node] == labels[next] {
				if nextCount == k {
					continue
				}

				nextCount++
			} else {
				nextCount = 1
			}

			if nextCost < dist[next][nextCount] {
				dist[next][nextCount] = nextCost

				heap.Push(&pq, &State{node: next, cost: nextCost, count: nextCount})
			}

		}
	}

	return -1
}