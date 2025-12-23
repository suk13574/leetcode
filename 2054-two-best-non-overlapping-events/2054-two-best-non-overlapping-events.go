import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)

	starts := make([][]int, n)
	ends := make([][]int, n)
	copy(starts, events)
	copy(ends, events)

	sort.Slice(starts, func(i, j int) bool { return starts[i][0] < starts[j][0] })
	sort.Slice(ends, func(i, j int) bool { return ends[i][1] < ends[j][1] })

	res := 0
	best := 0
	j := 0
	for i := 0; i < n; i++ {
		start := starts[i][0]
		value := starts[i][2]

		for j < n && ends[j][1] < start {
			best = max(best, ends[j][2])
			j++
		}

		res = max(res, value+best)
	}
	return res
}