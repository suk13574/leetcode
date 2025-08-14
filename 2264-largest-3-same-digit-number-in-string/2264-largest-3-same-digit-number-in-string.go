func largestGoodInteger(num string) string {
	var best byte
	n := len(num)

	for i := 0; i+2 < n; i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			if num[i] > best {
				best = num[i]
			}
			if best == '9' {
				break
			}
		}
	}

	if best == '0' {
		return ""
	}

	return string([]byte{best, best, best})
}