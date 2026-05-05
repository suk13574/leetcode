func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 1
	oldTail := head

	for oldTail.Next != nil {
		n++
		oldTail = oldTail.Next
	}

	r := k % n
	if r == 0 {
		return head
	}

	stepsToNewTail := n - 1 - r

	newTail := head
	for i := 0; i < stepsToNewTail; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next

	newTail.Next = nil
	oldTail.Next = head

	return newHead
}