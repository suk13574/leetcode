type Jump struct {
	idx  int
	jump int
}

func minJumps(arr []int) int {
	n := len(arr)

	sameNums := make(map[int][]int)

	for i, num := range arr {
		if _, ok := sameNums[num]; !ok {
			sameNums[num] = []int{}
		}

		sameNums[num] = append(sameNums[num], i)
	}

	visited := make([]bool, n)

	queue := []Jump{{0, 0}}
	visited[0] = true
	for cur := 0; cur < len(queue); cur++ {
		j := queue[cur]

		if j.idx == n-1 {
			return j.jump
		}

		for _, next := range []int{j.idx + 1, j.idx - 1} {
			if next >= 0 && next < n && !visited[next] {
				queue = append(queue, Jump{next, j.jump + 1})
				visited[next] = true
			}
		}

		for _, next := range sameNums[arr[j.idx]] {
			if !visited[next] {
				queue = append(queue, Jump{next, j.jump + 1})
				visited[next] = true
			}

			delete(sameNums, arr[j.idx])
		}
	}

	return -1
}