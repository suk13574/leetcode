/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func nodesBetweenCriticalPoints(head *ListNode) []int {
	points := findCriticalPoints(head)

	if len(points) <= 1 {
		return []int{-1, -1}
	}

	maxDistance := points[len(points)-1] - points[0]
	minDistance := maxDistance

	for i := 1; i < len(points); i++ {
		minDistance = min(minDistance, points[i]-points[i-1])
	}

	return []int{minDistance, maxDistance}
}

func findCriticalPoints(head *ListNode) []int {
	points := []int{}

	prev := head.Val
	head = head.Next
	idx := 1

	for head.Next != nil {
		now := head.Val
		next := head.Next.Val

		// maxima
		if prev > now && next > now {
			points = append(points, idx)
		}

		// minima
		if prev < now && next < now {
			points = append(points, idx)
		}

		idx++
		prev = now
		head = head.Next
	}

	return points
}