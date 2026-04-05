func judgeCircle(moves string) bool {
	td := 0
	lr := 0

	for _, m := range moves {
		switch m {
		case 'U':
			td++
		case 'D':
			td--
		case 'L':
			lr--
		case 'R':
			lr++
		}
	}

	if td == 0 && lr == 0 {
		return true
	}
	return false
}