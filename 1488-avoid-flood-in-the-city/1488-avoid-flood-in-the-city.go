import "sort"

func avoidFlood(rains []int) []int {
	n := len(rains)
	res := make([]int, n)

	lastRain := make(map[int]int)
	dryDays := []int{}

	for i, r := range rains {
		if r == 0 {
			res[i] = 1
			dryDays = append(dryDays, i)
			continue
		}

		res[i] = -1

		if last, ok := lastRain[r]; ok {
			j := sort.Search(len(dryDays), func(k int) bool { return dryDays[k] > last })
			if j == len(dryDays) {
				return []int{}
			}

			res[dryDays[j]] = r
			dryDays = append(dryDays[:j], dryDays[j+1:]...)
		}
		lastRain[r] = i
	}
	return res
}