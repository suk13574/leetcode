import (
	"math"
	"sort"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])

	res := make([][]int, r-k+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c-k+1)
	}

	findMinAbs := func(x, y int) int {
		sub := make([]int, 0, k*k)
		for i := 0; i < k; i++ {
			sub = append(sub, grid[x+i][y:y+k]...)
		}

		sort.Ints(sub)
		minGap := math.MaxInt

		for i := 1; i < len(sub); i++ {
			if sub[i] == sub[i-1] {
				continue
			}

			minGap = min(minGap, sub[i]-sub[i-1])
		}

		if minGap == math.MaxInt {
			return 0
		}
		return minGap
	}

	for i := 0; i+k-1 < r; i++ {
		for j := 0; j+k-1 < c; j++ {
			res[i][j] = findMinAbs(i, j)
		}
	}

	return res
}