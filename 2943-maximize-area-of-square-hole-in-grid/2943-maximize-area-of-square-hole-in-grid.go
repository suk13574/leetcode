func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)

	maxH := 1
	nowH := 1
	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i-1]+1 {
			nowH++
			continue
		}
		maxH = max(maxH, nowH)
		nowH = 1
	}
	maxH = max(maxH, nowH)

	maxV := 1
	nowV := 1
	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i-1]+1 {
			nowV++
			continue
		}
		maxV = max(maxV, nowV)
		nowV = 1
	}
	maxV = max(maxV, nowV)

	maxLen := 1 + min(maxH, maxV)

	return maxLen * maxLen
}