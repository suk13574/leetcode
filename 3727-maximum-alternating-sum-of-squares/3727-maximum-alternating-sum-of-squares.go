import "sort"

func maxAlternatingSum(nums []int) int64 {
    n := len(nums)

    for i := 0; i < n; i++ {
        nums[i] *= nums[i]
    }

    sort.Ints(nums)

    res := int64(0)
    
    left, right := 0, n-1
    for i := 0; i < n; i++ {
        if i%2 == 0 {
            res += int64(nums[right])
            right--
        } else {
            res -= int64(nums[left])
            left++
        }
    }

    return res
}