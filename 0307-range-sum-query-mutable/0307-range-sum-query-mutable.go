type SegmentTree struct {
	arr  []int
	tree []int
}

type NumArray struct {
	st *SegmentTree
}

func Constructor(nums []int) NumArray {
	st := SegmentTree{nums, make([]int, 4*len(nums))}
	st.build(0, len(nums)-1, 1)
	return NumArray{&st}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.st.arr[index]
	this.st.update(0, len(this.st.arr)-1, 1, index, diff)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.st.query(0, len(this.st.arr)-1, 1, left, right)
}

func (s *SegmentTree) build(start, end, node int) {
	if start == end {
		s.tree[node] = s.arr[start]
		return
	}

	mid := (start + end) / 2
	s.build(start, mid, node*2)
	s.build(mid+1, end, node*2+1)
	s.tree[node] = s.tree[node*2] + s.tree[node*2+1]
}

func (s *SegmentTree) update(start, end, node, i, diff int) {
	if i < start || end < i {
		return
	}

	s.tree[node] += diff

	if start != end {
		mid := (start + end) / 2
		s.update(start, mid, node*2, i, diff)
		s.update(mid+1, end, node*2+1, i, diff)
	}
}

func (s *SegmentTree) query(start, end, node, left, right int) int {
	if right < start || end < left {
		return 0
	}

	if left <= start && end <= right {
		return s.tree[node]
	}

	mid := (start + end) / 2

	leftSum := s.query(start, mid, node*2, left, right)
	rightSum := s.query(mid+1, end, node*2+1, left, right)

	return leftSum + rightSum
}