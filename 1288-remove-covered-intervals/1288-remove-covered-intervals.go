func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{}
	for _, interval := range intervals {
		r := interval[1]

		top := len(stack) - 1
		if top < 0 {
			stack = append(stack, interval)
			continue
		}

		topR := stack[top][1]

		if topR < r {
			stack = append(stack, interval)
		}
	}

	return len(stack)
}