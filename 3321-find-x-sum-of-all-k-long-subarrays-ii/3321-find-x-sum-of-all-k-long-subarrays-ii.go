import (
	"container/heap"
)

type entry struct {
	freq int
	val  int
	ver  int
}

/********** cold: MaxHeap (remaining entries) **********/
type coldHeap []entry

func (h coldHeap) Len() int { return len(h) }
func (h coldHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val > h[j].val
}
func (h coldHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *coldHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *coldHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

/********** hot: MinHeap (top-x by frequency) **********/
type hotHeap []entry

func (h hotHeap) Len() int { return len(h) }
func (h hotHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}
	return h[i].val < h[j].val
}
func (h hotHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hotHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *hotHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

/********** solver **********/
type solver struct {
	k, x int

	freq  map[int]int  // frequency of num
	ver   map[int]int  // version of num
	inHot map[int]bool // is num in hot?

	hot     hotHeap
	cold    coldHeap
	hotSize int
	sumHot  int64
}

func newSolver(k, x int) *solver {
	s := &solver{
		k:     k,
		x:     x,
		freq:  make(map[int]int),
		ver:   make(map[int]int),
		inHot: make(map[int]bool),
	}
	heap.Init(&s.hot)
	heap.Init(&s.cold)

	return s
}

func (s *solver) pushCurrent(val int) {
	if s.freq[val] == 0 {
		return
	}

	s.ver[val]++                                             // upgrade version
	e := entry{freq: s.freq[val], val: val, ver: s.ver[val]} // push new entry in heaps
	if s.inHot[val] {
		heap.Push(&s.hot, e)
	} else {
		heap.Push(&s.cold, e)
	}
}

func (s *solver) valid(e entry, expectHot bool) bool {
	if s.freq[e.val] == 0 {
		return false
	}
	if s.ver[e.val] != e.ver {
		return false
	}
	if s.freq[e.val] != e.freq {
		return false
	}
	if s.inHot[e.val] != expectHot {
		return false
	}
	return true
}

func (s *solver) cleanTopHot() {
	for s.hot.Len() > 0 {
		top := s.hot[0]
		if s.valid(top, true) {
			return
		}
		heap.Pop(&s.hot)
	}
}

func (s *solver) cleanTopCold() {
	for s.cold.Len() > 0 {
		top := s.cold[0]
		if s.valid(top, false) {
			return
		}
		heap.Pop(&s.cold)
	}
}

// cold -> hot
func (s *solver) promoteOne() bool {
	s.cleanTopCold()
	if s.cold.Len() == 0 {
		return false
	}

	e := heap.Pop(&s.cold).(entry)

	s.inHot[e.val] = true
	s.hotSize++
	s.sumHot += int64(e.val * e.freq)
	s.pushCurrent(e.val)

	return true
}

// hot -> cold
func (s *solver) demoteOne() bool {
	s.cleanTopHot()
	if s.hot.Len() == 0 {
		return false
	}

	e := heap.Pop(&s.hot).(entry)

	s.inHot[e.val] = false
	s.hotSize--
	s.sumHot -= int64(e.val * e.freq)
	s.pushCurrent(e.val)

	return true
}

func better(a, b entry) bool {
	if a.freq != b.freq {
		return a.freq > b.freq
	}
	return a.val > b.val
}

func (s *solver) rebalance() {
	// If hotSize is smaller than x, promote from cold until size == x.
	for s.hotSize < s.x {
		if !s.promoteOne() {
			break
		}
	}

	// If hotSize is larger than x, demote from hot until size == x.
	for s.hotSize > s.x {
		if !s.demoteOne() {
			break
		}
	}

	for {
		s.cleanTopCold()
		s.cleanTopHot()
		if s.cold.Len() == 0 || s.hot.Len() == 0 {
			break
		}
		cb := s.cold[0] // best in cold
		hw := s.hot[0]  // worst in hot
		if better(cb, hw) {
			// Swap boundary: demote worst from hot, promote best from cold.
			s.demoteOne()
			s.promoteOne()
		} else {
			break
		}
	}
}

func (s *solver) add(val int) {
	if s.freq[val] == 0 {
		s.inHot[val] = false
	}

	if s.inHot[val] {
		s.sumHot += int64(val)
	}

	s.freq[val]++
	s.pushCurrent(val)
	s.rebalance()
}

func (s *solver) remove(val int) {
	if s.freq[val] == 0 {
		return
	}

	if s.inHot[val] {
		s.sumHot -= int64(val)
	}
	s.freq[val]--
	s.pushCurrent(val)

	if s.freq[val] == 0 {
		if in, ok := s.inHot[val]; ok && in {
			s.hotSize--
		}
		delete(s.freq, val)
		delete(s.ver, val)
		delete(s.inHot, val)
	}
	s.rebalance()
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	s := newSolver(k, x)

	// init window
	for i := 0; i < k; i++ {
		s.add(nums[i])
	}

	res := make([]int64, 0, n-k+1)
	res = append(res, s.sumHot)

	for r := k; r < n; r++ {
		out := nums[r-k]
		in := nums[r]

		if out != in {
			s.remove(out)
			s.add(in)
		}

		res = append(res, s.sumHot)
	}
	return res
}