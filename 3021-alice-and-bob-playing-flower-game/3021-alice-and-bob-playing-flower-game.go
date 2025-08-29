func flowerGame(n int, m int) int64 {
	evenN, oddN := n/2, (n+1)/2
	evenM, oddM := m/2, (m+1)/2

	return int64(evenN*oddM + oddN*evenM)
}