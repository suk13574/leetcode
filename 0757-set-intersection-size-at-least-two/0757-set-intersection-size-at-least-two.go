import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	x, y := intervals[0][1]-1, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		val := intervals[i]
		if x >= val[0] {
			continue
		} else if y >= val[0] {
			x = y
			y = val[1]
			res += 1
		} else {
			x = val[1] - 1
			y = val[1]
			res += 2
		}
	}

	return res + 2
}