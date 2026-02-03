func isTrionic(nums []int) bool {
	n := len(nums)
	i := 0

    for i+1 < n && nums[i] < nums[i+1] {
        i++
    }
    p := i
    if p == 0 {
        return false
    }

    for i+1 < n && nums[i] > nums[i+1] {
        i++
    }
    q := i
    if p == q || q >= n-1 {
        return false
    }

    for i+1 < n && nums[i] < nums[i+1] {
        i++
    }

    return i == n-1
}