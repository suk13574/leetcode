/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	res := true

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		gap := left - right
		if gap < 0 {
			gap = -gap
		}
		if gap > 1 {
			res = false
		}

		return max(left, right) + 1
	}

	dfs(root)
	return res
}