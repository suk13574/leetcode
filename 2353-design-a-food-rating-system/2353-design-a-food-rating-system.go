// --- max heap ---
type node struct {
	food    string
	rating  int
	cuisine string
	idx     int
}

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].rating != pq[j].rating {
		return pq[i].rating > pq[j].rating
	}
	return pq[i].food < pq[j].food
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x any) {
	n := x.(*node)
	n.idx = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	x.idx = -1
	*pq = old[:n-1]
	return x
}

// --- FoodRatings ---
type FoodRatings struct {
	foodNode map[string]*node
	cuiPq    map[string]*PriorityQueue
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodNode: make(map[string]*node, len(foods)),
		cuiPq:    make(map[string]*PriorityQueue, len(ratings)),
	}

	for i := range foods {
		f, c, r := foods[i], cuisines[i], ratings[i]

		if _, ok := fr.cuiPq[c]; !ok {
			fr.cuiPq[c] = &PriorityQueue{}
			heap.Init(fr.cuiPq[c])
		}

		n := &node{food: f, rating: r, cuisine: c}
		fr.foodNode[f] = n
		heap.Push(fr.cuiPq[c], n)
	}
	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	fn := fr.foodNode[food]
	fn.rating = newRating
	heap.Fix(fr.cuiPq[fn.cuisine], fn.idx)
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	pq := fr.cuiPq[cuisine]
	return (*pq)[0].food
}