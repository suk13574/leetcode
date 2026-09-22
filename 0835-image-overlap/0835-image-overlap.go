func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)

	type point struct {
		r, c int
	}

	ones1 := []point{}
	ones2 := []point{}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if img1[r][c] == 1 {
				ones1 = append(ones1, point{r, c})
			}

			if img2[r][c] == 1 {
				ones2 = append(ones2, point{r, c})
			}
		}
	}

    moves := make(map[point]int)
    res := 0
    
    for _, one1 := range ones1 {
        for _, one2 := range ones2 {
            dr := one2.r - one1.r
            dc := one2.c - one1.c

            p := point{dr, dc}
            moves[p]++

            res = max(res, moves[p])
        }
    }

    return res
}