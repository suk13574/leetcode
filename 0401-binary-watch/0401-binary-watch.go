func readBinaryWatch(turnedOn int) []string {
	hw := []int{1, 2, 4, 8}
	mw := []int{1, 2, 4, 8, 16, 32}

	res := []string{}

	var dfs func(idx, cnt, hour, min int)
	dfs = func(idx, cnt, hour, min int) {
		if hour > 11 || min > 59 {
			return
		}

		if cnt == 0 {
			res = append(res, fmt.Sprintf("%d:%02d", hour, min))
			return
		}

		if idx == 10 {
			return
		}

		if idx < 4 {
			dfs(idx+1, cnt-1, hour+hw[idx], min)
		} else {
			dfs(idx+1, cnt-1, hour, min+mw[idx-4])
		}

		dfs(idx+1, cnt, hour, min)
	}

    dfs(0, turnedOn, 0, 0)
    
    return res
}