func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, o := range operations {
		if o == "++X" || o == "X++" {
			x++
		} else if o == "--X" || o == "X--" {
			x--
		}
	}
	return x
}