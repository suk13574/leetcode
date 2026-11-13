type Server struct {
	id      int
	weight  int
	endTime int64
}

type FreeServer []*Server

func (pq FreeServer) Len() int { return len(pq) }
func (pq FreeServer) Less(i, j int) bool {
	if pq[i].weight == pq[j].weight {
		return pq[i].id < pq[j].id
	}
	return pq[i].weight < pq[j].weight
}
func (pq FreeServer) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *FreeServer) Push(x any)   { *pq = append(*pq, x.(*Server)) }
func (pq *FreeServer) Pop() any {
	old := *pq
	n := len(old)
	server := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return server
}

type BusyServer []*Server

func (pq BusyServer) Len() int           { return len(pq) }
func (pq BusyServer) Less(i, j int) bool { return pq[i].endTime < pq[j].endTime }
func (pq BusyServer) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *BusyServer) Push(x any)        { *pq = append(*pq, x.(*Server)) }
func (pq *BusyServer) Pop() any {
	old := *pq
	n := len(old)
	server := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return server
}

func assignTasks(servers []int, tasks []int) []int {
	n := len(servers)

	free := make(FreeServer, n)

	for i := 0; i < n; i++ {
		free[i] = &Server{id: i, weight: servers[i]}
	}
	heap.Init(&free)

	busy := make(BusyServer, 0)
	heap.Init(&busy)

	res := make([]int, len(tasks))
	curTime := int64(0)

	for i, task := range tasks {
		curTime = max(curTime, int64(i))

		for busy.Len() > 0 && busy[0].endTime <= curTime {
			server := heap.Pop(&busy).(*Server)
			heap.Push(&free, server)
		}

		if free.Len() == 0 {
			curTime = busy[0].endTime

			for busy.Len() > 0 && busy[0].endTime <= curTime {
				server := heap.Pop(&busy).(*Server)
				heap.Push(&free, server)
			}
		}

		server := heap.Pop(&free).(*Server)
		server.endTime = curTime + int64(task)
		res[i] = server.id
		heap.Push(&busy, server)
	}

	return res
}