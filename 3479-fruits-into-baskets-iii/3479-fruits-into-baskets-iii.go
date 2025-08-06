func numOfUnplacedFruits(fruits []int, baskets []int) int {
	result := 0

	n := len(baskets)
	st := New(baskets)
	st.build(0, n-1, 1)

	for _, fruit := range fruits {
		idx := st.findFirst(0, n-1, 1, fruit)

		if idx == -1 {
			result++
		} else {
			st.update(0, n-1, 1, idx, -1)
		}
	}

	return result
}

type SegmentTree struct {
	arr  []int
	tree []int
}

func New(arr []int) *SegmentTree {
	return &SegmentTree{arr, make([]int, 4*len(arr))}
}

func (s *SegmentTree) build(start, end, node int) {
	if start == end {
		s.tree[node] = s.arr[start]
		return
	}
	mid := (start + end) / 2

	s.build(start, mid, 2*node)
	s.build(mid+1, end, 2*node+1)
	s.tree[node] = max(s.tree[2*node], s.tree[2*node+1])
}

func (s *SegmentTree) findFirst(start, end, node, fruit int) int {
	if s.tree[node] < fruit {
		return -1
	}

	if start == end {
		return start
	}

	mid := (start + end) / 2

	leftIdx := s.findFirst(start, mid, 2*node, fruit)
	if leftIdx != -1 {
		return leftIdx
	}

	return s.findFirst(mid+1, end, 2*node+1, fruit)
}

func (s *SegmentTree) update(start, end, node, i, newVal int) {
	if i < start || end < i {
		return
	}

	if start == end {
		s.tree[node] = newVal
		return
	}

	mid := (start + end) / 2
	s.update(start, mid, 2*node, i, newVal)
	s.update(mid+1, end, 2*node+1, i, newVal)

	s.tree[node] = max(s.tree[2*node], s.tree[2*node+1]) // update max value
}