import "sort"

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	const MOD = 1_000_000_007

	hLines := make([]int, 0, len(hFences)*2)
	hLines = append(hLines, 1, m)
	hLines = append(hLines, hFences...)

	vLines := make([]int, 0, len(vFences)*2)
	vLines = append(vLines, 1, n)
	vLines = append(vLines, vFences...)

	hSet := buildDiffSet(hLines)
	vSet := buildDiffSet(vLines)

	best := 0
	if len(hSet) <= len(vSet) {
		for d := range hSet {
			if _, ok := vSet[d]; ok && d > best {
				best = d
			}
		}
	} else {
		for d := range vSet {
			if _, ok := hSet[d]; ok && d > best {
				best = d
			}
		}
	}

	if best == 0 {
		return -1
	}

	area := (int64(best) * int64(best)) % MOD
	return int(area % MOD)
}

func buildDiffSet(lines []int) map[int]struct{} {
	sort.Ints(lines)

	n := len(lines)
	set := make(map[int]struct{}, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			set[lines[j]-lines[i]] = struct{}{}
		}
	}

	return set
}
