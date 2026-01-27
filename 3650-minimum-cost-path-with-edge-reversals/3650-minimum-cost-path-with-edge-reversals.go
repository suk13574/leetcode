import "container/heap"

type Edge struct {
	to   int
	cost int
}

type Item struct {
	v    int
	dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func minCost(n int, edges [][]int) int {
	graph := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{to: v, cost: w})
		graph[v] = append(graph[v], Edge{to: u, cost: w * 2})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1 << 31
	}
	dist[0] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{v: 0, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		v := cur.v
		if cur.dist != dist[v] {
			continue
		}
		if v == n-1 {
			return cur.dist
		}

		for _, e := range graph[v] {
			newDist := cur.dist + e.cost
			if newDist < dist[e.to] {
				dist[e.to] = newDist
				heap.Push(&pq, &Item{v: e.to, dist: newDist})
			}
		}
	}
	return -1
}