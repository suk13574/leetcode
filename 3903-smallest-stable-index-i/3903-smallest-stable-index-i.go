func firstStableIndex(nums []int, k int) int {
    n := len(nums)
    prefixMax := make([]int, n+1)
    suffixMin := make([]int, n+1)

    for i := 0; i < n; i++ {
        prefixMax[i+1] = max(prefixMax[i], nums[i])
    }

    suffixMin[n] = nums[n-1]
    for i := n-1; i >= 0; i-- {
        suffixMin[i] = min(suffixMin[i+1], nums[i])
    }

    for i := 0; i < n; i++ {
        if prefixMax[i+1] - suffixMin[i] <= k {
            return i
        }
    }

    return -1
}