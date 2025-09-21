import "container/heap"

// --- min heap ---
type Movie struct {
	shop  int
	movie int
	price int
	idx   int
}

// Unrented heap: price ↑, shop ↑
type UnrentHeap []*Movie

func (h UnrentHeap) Len() int { return len(h) }

func (h UnrentHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	return h[i].shop < h[j].shop
}

func (h UnrentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}

func (h *UnrentHeap) Push(x any) {
	m := x.(*Movie)
	m.idx = len(*h)
	*h = append(*h, m)
}

func (h *UnrentHeap) Pop() any {
	old := *h
	n := len(old)
	m := old[n-1]
	m.idx = -1
	*h = old[:n-1]
	return m
}

// Rented heap: price ↑, shop ↑, movie ↑
type RentHeap []*Movie

func (h RentHeap) Len() int { return len(h) }

func (h RentHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	if h[i].shop != h[j].shop {
		return h[i].shop < h[j].shop
	}
	return h[i].movie < h[j].movie
}

func (h RentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}

func (h *RentHeap) Push(x any) {
	m := x.(*Movie)
	m.idx = len(*h)
	*h = append(*h, m)
}

func (h *RentHeap) Pop() any {
	old := *h
	n := len(old)
	m := old[n-1]
	m.idx = -1
	*h = old[:n-1]
	return m
}

// --- MovieRentingSystem ---
type Key struct {
	shop  int
	movie int
}

type MovieRentingSystem struct {
	price    map[Key]int
	isRented map[Key]bool

	unrentPQ map[int]*UnrentHeap
	rentPQ   *RentHeap

	unrent map[Key]*Movie
	rent   map[Key]*Movie
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	r := &RentHeap{}
	heap.Init(r)

	mrs := MovieRentingSystem{
		price:    make(map[Key]int),
		isRented: make(map[Key]bool),
		unrentPQ: make(map[int]*UnrentHeap),
		rentPQ:   r,
		unrent:   make(map[Key]*Movie),
		rent:     make(map[Key]*Movie),
	}

	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]

		key := Key{shop: shop, movie: movie}
		mrs.price[key] = price
		mrs.isRented[key] = false

		if _, ok := mrs.unrentPQ[movie]; !ok {
			h := &UnrentHeap{}
			heap.Init(h)
			mrs.unrentPQ[movie] = h
		}

		m := &Movie{shop: shop, movie: movie, price: price}
		heap.Push(mrs.unrentPQ[movie], m)
		mrs.unrent[key] = m
	}

	return mrs
}

func (mrs *MovieRentingSystem) Search(movie int) []int {
	h, ok := mrs.unrentPQ[movie]
	if !ok || h == nil || h.Len() == 0 {
		return []int{}
	}

	res := make([]int, 0, 5)
	picked := make([]*Movie, 0, 5)

	for h.Len() > 0 && len(res) < 5 {
		top := heap.Pop(h).(*Movie)
		res = append(res, top.shop)
		picked = append(picked, top)
	}

	for _, m := range picked {
		heap.Push(h, m)
	}

	return res
}

func (mrs *MovieRentingSystem) Rent(shop int, movie int) {
	key := Key{shop: shop, movie: movie}

	// unrent -> remove
	m := mrs.unrent[key]

	h := mrs.unrentPQ[movie]
	heap.Remove(h, m.idx)

	delete(mrs.unrent, key)

	// rent -> add
	newMovie := &Movie{shop: shop, movie: movie, price: mrs.price[key]}
	heap.Push(mrs.rentPQ, newMovie)
	mrs.rent[key] = newMovie

	mrs.isRented[key] = true
}

func (mrs *MovieRentingSystem) Drop(shop int, movie int) {
	key := Key{shop: shop, movie: movie}

	// rent -> remove
	m := mrs.rent[key]

	h := mrs.rentPQ
	heap.Remove(h, m.idx)

	delete(mrs.rent, key)

	// unrent -> add
	newMovie := &Movie{shop: shop, movie: movie, price: mrs.price[key]}
	heap.Push(mrs.unrentPQ[movie], newMovie)
	mrs.unrent[key] = newMovie

	mrs.isRented[key] = false
}

func (mrs *MovieRentingSystem) Report() [][]int {
	h := mrs.rentPQ

	res := make([][]int, 0, 5)
	picked := make([]*Movie, 0, 5)

	for h.Len() > 0 && len(res) < 5 {
		top := heap.Pop(h).(*Movie)
		// report 포맷: [shop, movie]
		res = append(res, []int{top.shop, top.movie})
		picked = append(picked, top)
	}

	for _, m := range picked {
		heap.Push(h, m)
	}
	return res
}