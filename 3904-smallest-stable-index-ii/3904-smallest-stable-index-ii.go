func firstStableIndex(nums []int, k int) int {
    n := len(nums)

    suffixMin := make([]int, n)
    suffixMin[n-1] = nums[n-1]
    
    for i := n-2; i >= 0; i-- {
        suffixMin[i] = min(suffixMin[i+1], nums[i])
    }

    prefixMax := 0

    for i := 0; i < n; i++ {
        prefixMax = max(prefixMax, nums[i])
        
        if prefixMax - suffixMin[i] <= k {
            return i
        }
    }

    return -1
}