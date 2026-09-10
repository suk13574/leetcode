func averageOfSubtree(root *TreeNode) int {
	res := 0

	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		sumL, nL := dfs(node.Left)
		sumR, nR := dfs(node.Right)

		sum := node.Val + sumL + sumR
		count := 1 + nL + nR

		if sum/count == node.Val {
			res++
		}

		return sum, count
	}

	dfs(root)

	return res
}