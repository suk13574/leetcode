import "math/bits"

func countPalindromicSubsequence(s string) int {
	n := len(s)

	left := make([]int, n)
	right := make([]int, n)

	mask := 0
	for i := 0; i < n; i++ {
		left[i] = mask
		c := int(s[i] - 'a')
		mask |= 1 << c
	}

	mask = 0
	for i := n - 1; i >= 0; i-- {
		right[i] = mask
		c := int(s[i] - 'a')
		mask |= 1 << c
	}

	seen := make([]int, 26)
	res := 0

	for i := 1; i <= n-2; i++ {
		mid := int(s[i] - 'a')

		both := left[i] & right[i]

		newMask := both & ^seen[mid]
		if newMask == 0 {
			continue
		}

		res += bits.OnesCount(uint(newMask))
		seen[mid] |= newMask
	}

	return res
}
