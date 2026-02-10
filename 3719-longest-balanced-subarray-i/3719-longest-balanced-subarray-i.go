func longestBalanced(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		oddMap := make(map[int]struct{})
		evenMap := make(map[int]struct{})
		for j := i; j < len(nums); j++ {
			v := nums[j]

			if v%2 == 0 {
				evenMap[v] = struct{}{}
			} else {
				oddMap[v] = struct{}{}
			}

			if len(evenMap) == len(oddMap) {
				res = max(res, j-i+1)
			}
		}
	}

	return res
}