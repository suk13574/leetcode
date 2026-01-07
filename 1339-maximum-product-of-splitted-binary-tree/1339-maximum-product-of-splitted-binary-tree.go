/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
    const MOD int64 = 1_000_000_007

    sumValues := []int64{}

	var dfs func(n *TreeNode) int64
	dfs = func(n *TreeNode) int64 {
        if n == nil {
            return 0
        }

        sum := int64(n.Val) + dfs(n.Left) + dfs(n.Right)
        sumValues = append(sumValues, sum)
        return sum
	}

    total := dfs(root)
    res := int64(-1)

    for _, s := range sumValues {
        cand := s * (total-s)
        if cand > res {
            res = cand
        }
    }

	return int(res % MOD)
}