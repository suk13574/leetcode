func countCollisions(directions string) int {
	res := 0
	RCnt := 0
	prevStop := false
	for i := 0; i < len(directions); i++ {
		d := directions[i]

		switch d {
		case 'R':
			RCnt++
			prevStop = true

		case 'S':
			if RCnt > 0 {
				res += RCnt
				RCnt = 0
			}
			prevStop = true

		case 'L':
			if RCnt > 0 {
				res += RCnt + 1
				RCnt = 0
			} else if prevStop {
				res += 1
			} else {
				prevStop = false
			}
		}

	}
	return res
}