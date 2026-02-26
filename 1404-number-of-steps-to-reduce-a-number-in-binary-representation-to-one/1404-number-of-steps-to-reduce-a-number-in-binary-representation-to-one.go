func numSteps(s string) int {
	b := []byte(s)
	res := 0
	for len(b) > 1 {
		n := len(b)

		if b[n-1] == '1' {
			res++

			for i := n - 1; i >= 0; i-- {
				if b[i] == '0' {
					b[i] = '1'
					break
				}
				b[i] = '0'
			}
		}

		if b[0] == '0' {
			return res + len(b)
		}

		b = b[:n-1]
		res++
	}

	return res
}