func largestGoodInteger(num string) string {
	result := ""
	n := len(num)

	for left, right := 0, 2; left < n && right < n; {
		mid := left + 1

		if num[left] == num[mid] {
			if num[left] == num[right] {
				result = max(result, num[left:right+1])
				left += 3
				right += 3
			} else {
				left = right
				right += 2
			}
		} else {
			left = mid
			right += 1
		}
	}

	return result
}