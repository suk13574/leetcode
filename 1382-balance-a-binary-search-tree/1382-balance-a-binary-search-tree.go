/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func balanceBST(root *TreeNode) *TreeNode {
	// inorder
	arr := make([]int, 0)

	stack := make([]*TreeNode, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		arr = append(arr, node.Val)
		cur = node.Right
	}

	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := (l + r) / 2
		node := &TreeNode{Val: arr[mid]}
		node.Left = build(l, mid-1)
		node.Right = build(mid+1, r)
		return node
	}

	return build(0, len(arr)-1)
}