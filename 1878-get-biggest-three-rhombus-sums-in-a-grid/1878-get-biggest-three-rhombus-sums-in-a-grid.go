func getBiggestThree(grid [][]int) []int {
	r := len(grid)
	c := len(grid[0])

	diag1 := make([][]int, r+1)
	diag2 := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		diag1[i] = make([]int, c+2)
		diag2[i] = make([]int, c+2)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			diag1[i+1][j+1] = diag1[i][j] + grid[i][j]
			diag2[i+1][j+1] = diag2[i][j+2] + grid[i][j]
		}
	}

	getDiag1 := func(r1, c1, r2, c2 int) int {
		if r1 > r2 {
			r1, r2 = r2, r1
			c1, c2 = c2, c1
		}
		if r2-r1 != c2-c1 {
			return 0
		}

		return diag1[r2+1][c2+1] - diag1[r1][c1]
	}

	getDiag2 := func(r1, c1, r2, c2 int) int {
		if r1 > r2 {
			r1, r2 = r2, r1
			c1, c2 = c2, c1
		}
		if r2-r1 != c1-c2 {
			return 0
		}

		return diag2[r2+1][c2+1] - diag2[r1][c1+2]
	}

	seen := make(map[int]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			seen[grid[i][j]] = true

			maxK := min(min(i, r-1-i), min(j, c-1-j))
			for k := 1; k <= maxK; k++ {
				topR, topC := i-k, j
				rightR, rightC := i, j+k
				bottomR, bottomC := i+k, j
				leftR, leftC := i, j-k

				sumDiag1 := getDiag1(topR, topC, rightR, rightC) + getDiag1(leftR, leftC, bottomR, bottomC)
				sumDiag2 := getDiag2(topR, topC, leftR, leftC) + getDiag2(rightR, rightC, bottomR, bottomC)
				sum := sumDiag1 + sumDiag2 - grid[topR][topC] - grid[rightR][rightC] - grid[bottomR][bottomC] - grid[leftR][leftC]
				seen[sum] = true
			}
		}
	}

	sums := make([]int, 0, len(seen))
	for s := range seen {
		sums = append(sums, s)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

    if len(sums) < 3 {
        return sums
    }
	return sums[:3]
}