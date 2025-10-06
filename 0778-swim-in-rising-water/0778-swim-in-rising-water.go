import (
	"container/heap"
)

type swim struct {
	t    int
	r, c int
}

type Item struct {
	t     int
	r, c  int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].t < pq[j].t }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func swimInWater(grid [][]int) int {
	n := len(grid)

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	pq := make(PriorityQueue, 0, n*n)
	heap.Init(&pq)

	heap.Push(&pq, &Item{t: grid[0][0], r: 0, c: 0})

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*Item)
		r, c, t := i.r, i.c, i.t

		if visited[r][c] {
			continue
		}
		visited[r][c] = true

		if r == n-1 && c == n-1 {
			return t
		}

		for _, m := range moves {
			nr, nc := r+m[0], c+m[1]
			if nr < 0 || nc < 0 || nr >= n || nc >= n {
				continue
			}
			if visited[nr][nc] {
				continue
			}

			nt := max(t, grid[nr][nc])
			heap.Push(&pq, &Item{t: nt, r: nr, c: nc})
		}
	}
	return 0
}