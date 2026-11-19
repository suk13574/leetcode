func firstMissingPositive(nums []int) int {
    n := len(nums)

    for i := 0; i < n; i++ {
        v := nums[i]
        
        for (0 < v && v <= n) && v != i+1 && nums[v-1] != v {
            nums[i], nums[v-1] = nums[v-1], nums[i]
            v = nums[i]
        }
    }

    idx := 0
    for ; idx < n; idx++ {
        if idx+1 != nums[idx] {
            break
        }
    }

    return idx+1
}