/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    siblingSum := make(map[int]int64)

    var dfs func(level int, n *TreeNode)
    dfs = func(level int, n *TreeNode) {
        if n.Left != nil {
            dfs(level+1, n.Left)
        }
        if n.Right != nil {
            dfs(level+1, n.Right)
        }
        siblingSum[level] += int64(n.Val)
    }

    dfs(1, root)

    maxVal := siblingSum[1]
    res := 1
    for k, v := range siblingSum {
        if v > maxVal {
            maxVal = v
            res = k
        }
        if v == maxVal {
            res = min(res, k)
        }
    }

    return res
}