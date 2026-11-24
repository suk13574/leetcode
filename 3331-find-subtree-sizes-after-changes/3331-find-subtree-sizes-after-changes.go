func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	tree := make([][]int, n)

	for node, p := range parent {
		if p == -1 {
			continue
		}

		tree[p] = append(tree[p], node)
	}

	newParent := make([]int, n)
	copy(newParent, parent)

	lastSeen := make([][]int, 26)

	var reparentDFS func(int)
	reparentDFS = func(node int) {
		strIdx := s[node] - 'a'

		ancestors := lastSeen[strIdx]

		if len(ancestors) > 0 {
			newParent[node] = ancestors[len(ancestors)-1]
		}

		lastSeen[strIdx] = append(lastSeen[strIdx], node)

		for _, child := range tree[node] {
			reparentDFS(child)
		}

		lastSeen[strIdx] = lastSeen[strIdx][:len(lastSeen[strIdx])-1]
	}

	reparentDFS(0)

	newTree := make([][]int, n)

	for node, p := range newParent {
		if p == -1 {
			continue
		}

		newTree[p] = append(newTree[p], node)
	}

	subs := make([]int, n)

	var subtreeDFS func(int)
	subtreeDFS = func(node int) {
		sub := 1

		for _, child := range newTree[node] {
			subtreeDFS(child)
			sub += subs[child]
		}

		subs[node] = sub
	}

	subtreeDFS(0)

	return subs
}