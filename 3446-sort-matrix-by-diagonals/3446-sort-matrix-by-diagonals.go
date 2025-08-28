import (
	"container/heap"
)

type Heap struct {
	data []int
	less func(a, b int) bool
}

func NewMinHeap() *Heap {
	return &Heap{less: func(a, b int) bool { return a < b }}
}
func NewMaxHeap() *Heap {
	return &Heap{less: func(a, b int) bool { return a > b }}
}

func (h Heap) Len() int { return len(h.data) }

func (h Heap) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }

func (h Heap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *Heap) Pop() any {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

func (h *Heap) PushInt(x int) {
	heap.Push(h, x)
}

func (h *Heap) PopInt() int {
	return heap.Pop(h).(int)
}

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	minHeap := NewMinHeap()
	heap.Init(minHeap)

	maxHeap := NewMaxHeap()
	heap.Init(maxHeap)

	for i := 0; i < n; i++ {
		for x, y := i, 0; x < n && y < n; {
			maxHeap.PushInt(grid[x][y])
			x++
			y++
		}
		for x, y := i, 0; x < n && y < n; {
			result[x][y] = maxHeap.PopInt()
			x++
			y++
		}
	}

	for j := 1; j < n; j++ {
		for x, y := 0, j; x < n && y < n; {
			minHeap.PushInt(grid[x][y])
			x++
			y++
		}
		for x, y := 0, j; x < n && y < n; {
			result[x][y] = minHeap.PopInt()
			x++
			y++
		}
	}
	return result
}