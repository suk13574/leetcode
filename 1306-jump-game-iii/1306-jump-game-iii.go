func canReach(arr []int, start int) bool {
	n := len(arr)

	queue := []int{}
	visited := make([]bool, n)

	queue = append(queue, start)

	for cur := 0; cur < len(queue); cur++ {
		idx := queue[cur]

		if arr[idx] == 0 {
			return true
		}

		if visited[idx] {
			continue
		}
		visited[idx] = true

		jump := arr[idx]

		next := idx + jump
		if next < n && !visited[next] {
			queue = append(queue, next)
		}

		next = idx - jump
		if next >= 0 && !visited[next] {
			queue = append(queue, next)
		}
	}

	return false
}