func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{
		{0, 1},  // north
		{1, 0},  // east
		{0, -1}, // south
		{-1, 0}, // west
	}

	obs := make(map[[2]int]bool)
	for _, o := range obstacles {
		obs[[2]int{o[0], o[1]}] = true
	}

	x, y := 0, 0
	dir := 0
	res := 0

	for _, c := range commands {
		if c == -1 {
			dir = (dir + 1) % 4
		} else if c == -2 {
			dir = (dir + 3) % 4
		} else {
			for step := 0; step < c; step++ {
				nx := x + dirs[dir][0]
				ny := y + dirs[dir][1]

				if obs[[2]int{nx, ny}] {
					break
				}

				x, y = nx, ny
				res = max(res, x*x+y*y)
			}
		}
	}

	return res
}