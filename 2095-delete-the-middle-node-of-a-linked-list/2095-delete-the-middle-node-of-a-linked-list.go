/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    prev := head
    middle := head
    tail := head

    for tail != nil && tail.Next != nil {
        prev = middle
        middle = middle.Next
        tail = tail.Next.Next
    }

    if prev == middle {
        return nil
    }

    prev.Next = middle.Next

    return head
}