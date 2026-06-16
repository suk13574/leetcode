func processStr(s string) string {
	res := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '#':
			res = duplicate(res)
		case '%':
			reverse(res)
		case '*':
			res = remove(res)
		default:
			res = append(res, ch)
		}
	}

	return string(res)
}

func duplicate(res []byte) []byte {
	return append(res, res...)
}

func reverse(res []byte) {
	n := len(res)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
}

func remove(res []byte) []byte {
	n := len(res)
	if n == 0 {
		return res
	}

	return res[:n-1]
}