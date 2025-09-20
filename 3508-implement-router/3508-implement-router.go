import "sort"

type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	memoryLimit int
	q           []Packet
	seen        map[Packet]struct{}
	destTimes   map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		q:           make([]Packet, 0, memoryLimit),
		seen:        make(map[Packet]struct{}),
		destTimes:   make(map[int][]int),
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	p := Packet{source: source, destination: destination, timestamp: timestamp}
	if _, ok := r.seen[p]; ok {
		return false
	}

	if len(r.q) == r.memoryLimit {
		oldest := r.q[0]
		delete(r.seen, oldest)
		r.q = r.q[1:]
		r.destTimes[oldest.destination] = r.destTimes[oldest.destination][1:]
	}

	r.seen[p] = struct{}{}
	r.q = append(r.q, p)
	r.destTimes[destination] = append(r.destTimes[destination], timestamp)
	return true
}

func (r *Router) ForwardPacket() []int {
	if len(r.q) == 0 {
		return []int{}
	}
	first := r.q[0]
	delete(r.seen, first)
	r.q = r.q[1:]
	r.destTimes[first.destination] = r.destTimes[first.destination][1:]

	return []int{first.source, first.destination, first.timestamp}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	ts := r.destTimes[destination]
	if len(ts) == 0 {
		return 0
	}

	left := sort.Search(len(ts), func(i int) bool { return ts[i] >= startTime })
	right := sort.Search(len(ts), func(i int) bool { return ts[i] > endTime })

	return right - left
}