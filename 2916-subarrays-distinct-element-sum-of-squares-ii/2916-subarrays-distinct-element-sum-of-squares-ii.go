const MOD int64 = 1_000_000_007

type Node struct {
	sum   int64
	sumSq int64
	lazy  int64
}

type SegmentTree struct {
	tree []Node
	n    int
}

func NewSegmentTree(n int) *SegmentTree {
	return &SegmentTree{
		tree: make([]Node, 4*n),
		n:    n,
	}
}

func (st *SegmentTree) apply(node, l, r int, k int64) {
	length := int64(r - l + 1)
	cur := &st.tree[node]

	oldSum := cur.sum

	add := (2*k*oldSum + length*k*k) % MOD

	cur.sumSq = (cur.sumSq + add) % MOD
	cur.sum = (cur.sum + length*k) % MOD
	cur.lazy = (cur.lazy + k) % MOD
}

func (st *SegmentTree) pushDown(node, l, r int) {
	if l == r {
		return
	}

	cur := &st.tree[node]

	if cur.lazy == 0 {
		return
	}

	mid := l + (r-l)/2

	st.apply(node*2, l, mid, cur.lazy)

	st.apply(node*2+1, mid+1, r, cur.lazy)

	cur.lazy = 0
}

func (st *SegmentTree) pullUp(node int) {
	left := node * 2
	right := node*2 + 1

	st.tree[node].sum = (st.tree[left].sum + st.tree[right].sum) % MOD

	st.tree[node].sumSq = (st.tree[left].sumSq + st.tree[right].sumSq) % MOD
}

func (st *SegmentTree) update(node, l, r, ql, qr int) {
	if r < ql || l > qr {
		return
	}

	if ql <= l && r <= qr {
		st.apply(node, l, r, 1)
		return
	}

	st.pushDown(node, l, r)

	mid := l + (r-l)/2

	st.update(node*2, l, mid, ql, qr)
	st.update(node*2+1, mid+1, r, ql, qr)

	st.pullUp(node)
}

func sumCounts(nums []int) int {
	n := len(nums)
	st := NewSegmentTree(n)

	last := make(map[int]int)
	var res int64

	for r, num := range nums {
		prev, ok := last[num]
		if !ok {
			prev = -1
		}

		st.update(1, 0, n-1, prev+1, r)

		res = (res + st.tree[1].sumSq) % MOD

		last[num] = r
	}

	return int(res)
}