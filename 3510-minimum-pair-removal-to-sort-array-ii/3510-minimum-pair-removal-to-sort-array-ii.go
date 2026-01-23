import "container/heap"

type PairSum struct {
	sum         int64
	left, right int
}

type MinHeap []*PairSum

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	if h[i].sum == h[j].sum {
		return h[i].left < h[j].left
	}
	return h[i].sum < h[j].sum
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	item := x.(*PairSum)
	*h = append(*h, item)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	ps := old[n-1]
	*h = old[:n-1]
	return ps
}

func minimumPairRemoval(nums []int) int {
	n := len(nums)
	val := make([]int64, n)
	prev := make([]int, n)
	next := make([]int, n)
	alive := make([]bool, n)

	minHeap := MinHeap{}
	heap.Init(&minHeap)

	for i := 0; i < n; i++ {
		val[i] = int64(nums[i])
		alive[i] = true

		if i == 0 {
			prev[i] = -1
		} else {
			prev[i] = i - 1
			sum := val[i-1] + val[i]
			heap.Push(&minHeap, &PairSum{sum: sum, left: i - 1, right: i})
		}

		if i == n-1 {
			next[i] = -1
		} else {
			next[i] = i + 1

		}
	}

	bad := 0
	for i := 0; i != -1; i = next[i] {
		j := next[i]
		if j == -1 {
			break
		}
		if val[i] > val[j] {
			bad++
		}
	}
	if bad == 0 {
		return 0
	}

	popValid := func() *PairSum {
		for {
			ps := heap.Pop(&minHeap).(*PairSum)
			if alive[ps.left] && alive[ps.right] &&
				next[ps.left] == ps.right &&
				ps.sum == val[ps.left]+val[ps.right] {
				return ps
			}
		}
	}

	decBad := func(l, r int) {
		if l == -1 || r == -1 {
			return
		}
		if val[l] > val[r] {
			bad--
		}
	}

	incBad := func(l, r int) {
		if l == -1 || r == -1 {
			return
		}
		if val[l] > val[r] {
			bad++
		}
	}

	res := 0
	for bad > 0 {
		res++

		ps := popValid()
		l := (*ps).left
		r := (*ps).right

		ll := prev[l]
		rr := next[r]

		// operation bad for old edges
		decBad(ll, l)
		decBad(l, r)
		decBad(r, rr)

		// merge
		val[l] += val[r]
		alive[r] = false

		next[l] = rr
		if rr != -1 {
			prev[rr] = l
		}

		if ll != -1 {
			heap.Push(&minHeap, &PairSum{sum: val[ll] + val[l], left: ll, right: l})
		}
		if rr != -1 {
			heap.Push(&minHeap, &PairSum{sum: val[l] + val[rr], left: l, right: rr})
		}

		//opertaion bad for new edges
		incBad(ll, l)
		incBad(l, rr)
	}

	return res
}