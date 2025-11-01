func modifiedList(nums []int, head *ListNode) *ListNode {
	del := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		del[v] = struct{}{}
	}

	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if _, ok := del[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}