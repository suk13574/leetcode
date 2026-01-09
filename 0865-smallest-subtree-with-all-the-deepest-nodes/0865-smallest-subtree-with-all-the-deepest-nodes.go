/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var dfs func(n *TreeNode) (*TreeNode, int)
	dfs = func(n *TreeNode) (*TreeNode, int) {
		if n == nil {
			return nil, 0
		}

		ln, ld := dfs(n.Left)
		rn, rd := dfs(n.Right)

		if ld == rd {
			return n, ld + 1
		}

		if ld > rd {
			return ln, ld + 1
		}
		return rn, rd + 1
	}

	res, _ := dfs(root)

	return res
}