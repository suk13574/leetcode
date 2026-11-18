/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	queue := []*TreeNode{root}

	levels := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			cur := queue[i]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		queue = queue[size:]
		levels = append(levels, level)
	}

	res := 0

	for _, level := range levels {
		n := len(level)

		if n <= 1 {
			continue
		}

		sorted := append([]int(nil), level...)
		sort.Ints(sorted)

		sortedIdx := make(map[int]int)
		for i, v := range sorted {
			sortedIdx[v] = i
		}

		targetIdx := make([]int, n)
		for i := 0; i < n; i++ {
			targetIdx[i] = sortedIdx[level[i]]
		}

		visited := make([]bool, n)
		for i := 0; i < n; i++ {
            if visited[i] {
                continue
            }

			cnt := 1
            visited[i] = true
            nextIdx := targetIdx[i]

			for !visited[nextIdx] {
				visited[nextIdx] = true
				cnt++
				nextIdx = targetIdx[nextIdx]
			}

            res += cnt - 1
		}
	}

	return res
}