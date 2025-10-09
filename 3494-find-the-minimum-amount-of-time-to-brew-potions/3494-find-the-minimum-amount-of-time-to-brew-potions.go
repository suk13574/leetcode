func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)

	before := make([]int64, n)
	after := make([]int64, n)
	var s int64
	for i := 0; i < n; i++ {
		before[i] = s
		s += int64(skill[i])
		after[i] = s
	}

	startTime := make([]int64, n)

	var startJ int64
	for j := 0; j < m; j++ {
		nowMana := int64(mana[j])

		need := startTime[0]
		for i := 1; i < n; i++ {
			v := startTime[i] - nowMana*before[i]
			need = max(need, v)
		}

		startJ = need

		for i := 0; i < n; i++ {
			startTime[i] = startJ + nowMana*after[i]
		}
	}

	return startTime[n-1]
}