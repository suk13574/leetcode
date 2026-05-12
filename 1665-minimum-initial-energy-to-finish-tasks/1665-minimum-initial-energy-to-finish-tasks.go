import "sort"

func minimumEffort(tasks [][]int) int {
	totalActual := 0
	maxMinimum := 0

	for _, task := range tasks {
		totalActual += task[0]
		maxMinimum = max(maxMinimum, task[1])
	}

	sort.Slice(tasks, func(i, j int) bool {
		gapI := tasks[i][1] - tasks[i][0]
		gapJ := tasks[j][1] - tasks[j][0]
		return gapI > gapJ
	})

	isEnoughEnergy := func(e int) bool {
		energy := e

		for _, task := range tasks {
			actual, minimum := task[0], task[1]

			if energy < minimum {
				return false
			}
			energy -= actual
		}
		return true
	}

	l := 0
	r := totalActual + maxMinimum

	for l < r {
		mid := l + (r-l)/2

		if isEnoughEnergy(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}