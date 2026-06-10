import (
	"container/heap"
	"math"
)

// === Segment Tree ===
type SegmentTree struct {
	n       int
	maxTree []int
	minTree []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)

	maxTree := make([]int, 2*n)
	minTree := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		maxTree[i] = math.MinInt
		minTree[i] = math.MaxInt
	}

	// leat node
	for i := 0; i < n; i++ {
		maxTree[n+i] = nums[i]
		minTree[n+i] = nums[i]
	}

	// parent node
	for i := n - 1; i > 0; i-- {
		maxTree[i] = max(maxTree[i*2], maxTree[i*2+1])
		minTree[i] = min(minTree[i*2], minTree[i*2+1])
	}

	return &SegmentTree{
		n:       n,
		maxTree: maxTree,
		minTree: minTree,
	}
}

func (st *SegmentTree) QueryMax(left, right int) int {
	left += st.n
	right += st.n

	res := math.MinInt

	for left <= right {
		if left%2 == 1 {
			res = max(res, st.maxTree[left])
			left++
		}

		if right%2 == 0 {
			res = max(res, st.maxTree[right])
			right--
		}
		left /= 2
		right /= 2
	}

	return res
}

func (st *SegmentTree) QueryMin(left, right int) int {
	left += st.n
	right += st.n

	res := math.MaxInt

	for left <= right {
		if left%2 == 1 {
			res = min(res, st.minTree[left])
			left++
		}

		if right%2 == 0 {
			res = min(res, st.minTree[right])
			right--
		}

		left /= 2
		right /= 2
	}

	return res
}

func (st *SegmentTree) Value(left, right int) int64 {
	return int64(st.QueryMax(left, right) - st.QueryMin(left, right))
}

// Priority Queue
type Item struct {
	value int64
	left  int
	right int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value > pq[j].value }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func maxTotalValue(nums []int, k int) int64 {
	st := NewSegmentTree(nums)

	pq := &PriorityQueue{}
	heap.Init(pq)

	n := len(nums)

	for left := 0; left < n; left++ {
		right := n - 1
		value := st.Value(left, right)

		heap.Push(pq, &Item{
			value: value,
			left:  left,
			right: right,
		})
	}

	var res int64 = 0

	for i := 0; i < k; i++ {
		item := heap.Pop(pq).(*Item)

		res += item.value

		nextRight := item.right - 1

		if nextRight >= item.left {
			nextValue := st.Value(item.left, nextRight)

			heap.Push(pq, &Item{
				value: nextValue,
				left:  item.left,
				right: nextRight,
			})
		}
	}

	return res
}