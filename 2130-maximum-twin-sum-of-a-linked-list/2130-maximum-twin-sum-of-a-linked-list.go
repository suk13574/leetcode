/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func pairSum(head *ListNode) int {
	half := head
	tail := head

	for tail != nil && tail.Next != nil {
		half = half.Next
		tail = tail.Next.Next
	}

	var prev *ListNode
	cur := half
	for cur != nil {
		nextNode := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextNode
	}

	left := head
	right := prev
	res := 0

	for right != nil {
		res = max(res, left.Val+right.Val)
		left = left.Next
		right = right.Next
	}

	return res
}