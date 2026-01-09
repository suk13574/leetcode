/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var res *TreeNode
	maxDepth := -1

	var dfs func(n *TreeNode, depth int) int
	dfs = func(n *TreeNode, depth int) int {
		if n == nil {
            maxDepth = max(maxDepth, depth-1)
			return depth-1
		}

        leftMax := dfs(n.Left, depth+1)
        rightMax := dfs(n.Right, depth+1)

        if leftMax == maxDepth && rightMax == maxDepth {
            res = n
        }

        if leftMax == maxDepth || rightMax == maxDepth {
            return maxDepth
        }

        return depth
	}

	dfs(root, 0)

	return res
}