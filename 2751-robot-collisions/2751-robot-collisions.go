type robot struct {
	pos    int
	health int
	dir    byte
	idx    int
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	robots := make([]robot, n)

	for i := 0; i < n; i++ {
		robots[i] = robot{pos: positions[i], health: healths[i], dir: directions[i], idx: i}
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].pos < robots[j].pos
	})

	stack := []int{}
	for i := 0; i < n; i++ {
		if robots[i].dir == 'R' {
			stack = append(stack, i)
			continue
		}

		// robots[i].dir == 'L'
		for len(stack) > 0 && robots[i].health > 0 {
			top := stack[len(stack)-1]

			if robots[top].health < robots[i].health {
				robots[i].health--
				robots[top].health = 0
				stack = stack[:len(stack)-1]
			} else if robots[top].health > robots[i].health {
				robots[top].health--
				robots[i].health = 0
			} else {
				robots[top].health = 0
				robots[i].health = 0
				stack = stack[:len(stack)-1]
			}
		}
	}

	survivors := make([]int, n)
	for _, r := range robots {
		if r.health > 0 {
			survivors[r.idx] = r.health
		}
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if survivors[i] > 0 {
			res = append(res, survivors[i])
		}
	}

	return res
}