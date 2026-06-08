/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	childSet := make(map[int]bool)

	for _, des := range descriptions {
		parent, child, isLeft := des[0], des[1], des[2]

		if _, ok := nodeMap[parent]; !ok {
			nodeMap[parent] = &TreeNode{Val: parent}
		}

		if _, ok := nodeMap[child]; !ok {
			nodeMap[child] = &TreeNode{Val: child}
		}

		pNode := nodeMap[parent]
		cNode := nodeMap[child]

		if isLeft == 1 {
			pNode.Left = cNode
		} else {
			pNode.Right = cNode
		}
		childSet[child] = true
	}

	for val, node := range nodeMap {
		if !childSet[val] {
			return node
		}
	}

	return nil
}