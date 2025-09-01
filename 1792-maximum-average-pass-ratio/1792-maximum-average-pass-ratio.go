type Classes struct {
	gain  float64
	pass  int
	total int
}

type PriorityQueue []*Classes

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].gain > pq[j].gain
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	classes := x.(*Classes)
	*pq = append(*pq, classes)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	delta := func(p, t int) float64 {
		return (float64(p+1) / float64(t+1)) - (float64(p) / float64(t))
	}

	n := len(classes)
	pq := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		p := classes[i][0]
		t := classes[i][1]
		pq[i] = &Classes{
			gain:  delta(p, t),
			pass:  p,
			total: t,
		}
	}
	heap.Init(&pq)

	for i := 0; i < extraStudents; i++ {
		class := heap.Pop(&pq).(*Classes)
		class.pass++
		class.total++
		class.gain = delta(class.pass, class.total)

		heap.Push(&pq, class)
	}

	var sum float64
	for pq.Len() != 0 {
		class := heap.Pop(&pq).(*Classes)
		sum += float64(class.pass) / float64(class.total)
	}
	return sum / float64(n)
}