func numberOfBeams(bank []string) int {
	res := 0
	prevCnt := 0
	for _, b := range bank {
		cnt := 0
		for i := 0; i < len(b); i++ {
			if b[i] == '1' {
				cnt++
			}
		}

		if cnt > 0 {
			res += prevCnt * cnt
			prevCnt = cnt
		}

		if cnt > 0 {

		}
	}

	return res
}