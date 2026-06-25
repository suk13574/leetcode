func countMajoritySubarrays(nums []int, target int) int {
    n := len(nums)
    res := 0
    for left := 0; left < n; left++ {
        sum := 0
        for right := left; right < n; right++ {
            if nums[right] == target {
                sum++
            } else {
                sum--
            }

            if sum > 0 {
                res++
            }
        }
    }
    
    return res
}
