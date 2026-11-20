import "sort"

func maxSelectedElements(nums []int) int {
	sort.Ints(nums)

	dp := make(map[int]int)
	res := 0

	for _, v := range nums {
		dp[v+1] = max(dp[v+1], dp[v]+1)
		dp[v] = max(dp[v], dp[v-1]+1)

		res = max(res, max(dp[v], dp[v+1]))
	}

	return res
}