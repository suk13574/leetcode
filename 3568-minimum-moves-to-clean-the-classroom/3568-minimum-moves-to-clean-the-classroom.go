type Node struct {
	r, c   int
	energy int
	mask   int
	moves  int
}

func minMoves(classroom []string, energy int) int {
	m, n := len(classroom), len(classroom[0])

	litterID := make([][]int, m)
	for r := range classroom {
		litterID[r] = make([]int, n)
		for c := range litterID[r] {
			litterID[r][c] = -1
		}
	}

	sr, sc := -1, -1
	litterCount := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch classroom[r][c] {
			case 'S':
				sr, sc = r, c

			case 'L':
				litterID[r][c] = litterCount
				litterCount++
			}
		}
	}

	targetMask := (1 << litterCount) - 1

	if targetMask == 0 {
		return 0
	}

	queue := []Node{{sr, sc, energy, 0, 0}}

	best := make([][][]int, m)
	for r := 0; r < m; r++ {
		best[r] = make([][]int, n)
		for c := 0; c < n; c++ {
			best[r][c] = make([]int, 1<<litterCount)
			for mask := range best[r][c] {
				best[r][c][mask] = -1
			}
		}
	}

	best[sr][sc][0] = energy

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for pos := 0; pos < len(queue); pos++ {
		cur := queue[pos]

		if cur.mask == targetMask {
			return cur.moves
		}

		if cur.energy == 0 {
			continue
		}

		for _, d := range dirs {
			nr := cur.r + d[0]
			nc := cur.c + d[1]

			if nr < 0 || nc < 0 || nr >= m || nc >= n {
				continue
			}

			nextEnergy := cur.energy - 1
			nextMask := cur.mask

			switch classroom[nr][nc] {
			case 'X':
				continue
			case 'R':
				nextEnergy = energy
			case 'L':
				id := litterID[nr][nc]
				nextMask |= 1 << id
			}

			if best[nr][nc][nextMask] >= nextEnergy {
				continue
			}

			best[nr][nc][nextMask] = nextEnergy
			queue = append(queue, Node{nr, nc, nextEnergy, nextMask, cur.moves + 1})
		}
	}

	return -1
}