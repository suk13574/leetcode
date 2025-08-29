func flowerGame(n int, m int) int64 {
	evenN, oddN := n/2, n/2
	evenM, oddM := m/2, m/2

	if n%2 == 1 {
		oddN++
	}
	if m%2 == 1 {
		oddM++
	}

	return (int64(evenN) * int64(oddM)) + int64(oddN)*int64(evenM)

}