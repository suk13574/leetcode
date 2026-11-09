import "container/heap"

type Server struct {
	id      int
	endTime int
	key     int
}

type BusyHeap []*Server

func (pq BusyHeap) Len() int           { return len(pq) }
func (pq BusyHeap) Less(i, j int) bool { return pq[i].endTime < pq[j].endTime }
func (pq BusyHeap) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *BusyHeap) Push(x any)        { *pq = append(*pq, x.(*Server)) }
func (pq *BusyHeap) Pop() any {
	old := *pq
	n := len(old)
	server := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return server
}

type FreeHeap []*Server

func (pq FreeHeap) Len() int           { return len(pq) }
func (pq FreeHeap) Less(i, j int) bool { return pq[i].key < pq[j].key }
func (pq FreeHeap) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *FreeHeap) Push(x any)        { *pq = append(*pq, x.(*Server)) }
func (pq *FreeHeap) Pop() any {
	old := *pq
	n := len(old)
	server := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return server
}

func busiestServers(k int, arrival []int, load []int) []int {
	n := len(arrival)

	free := make(FreeHeap, k)
	for i := 0; i < k; i++ {
		free[i] = &Server{id: i, key: i}
	}
	heap.Init(&free)

	busy := make(BusyHeap, 0)
	heap.Init(&busy)

	serverCnt := make([]int, k)

	for i := 0; i < n; i++ {
		currentTime := arrival[i]

		for busy.Len() > 0 && busy[0].endTime <= currentTime {
			server := heap.Pop(&busy).(*Server)

			offset := (server.id - i%k + k) % k
			server.key = i + offset

			heap.Push(&free, server)
		}

		if free.Len() == 0 {
			continue
		}

		server := heap.Pop(&free).(*Server)
		server.endTime = currentTime + load[i]
		serverCnt[server.id]++

		heap.Push(&busy, server)
	}

	maxCnt := 0
	for _, cnt := range serverCnt {
		maxCnt = max(maxCnt, cnt)
	}

	res := []int{}
	for i, cnt := range serverCnt {
		if cnt == maxCnt {
			res = append(res, i)
		}
	}

	return res
}