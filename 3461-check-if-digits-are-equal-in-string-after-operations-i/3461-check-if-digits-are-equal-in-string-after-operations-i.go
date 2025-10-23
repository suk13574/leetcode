func hasSameDigits(s string) bool {
	const MOD byte = 10
	ss := s
	for len(ss) > 2 {
		operS := make([]byte, len(ss)-1)
		for i := 0; i+1 < len(ss); i++ {
			ls := byte(ss[i] - '0')
			rs := byte(ss[i+1] - '0')
			operS[i] = '0' + (ls+rs)%MOD
		}
		ss = string(operS)
		println("ss = ", ss)
	}

	return ss[0] == ss[1]
}