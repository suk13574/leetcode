func rotateElements(nums []int, k int) []int {

    pos := []int{}

    for _, num := range nums {
        if num >= 0 {
            pos = append(pos, num)
        }
    }

    if len(pos) == 0 {
        return nums
    }
    
    k %= len(pos)

    if k > 0 {
        pos = append(pos[k:], pos[:k]...)
    }
    
    posIdx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            continue
        }

        nums[i] = pos[posIdx]
        posIdx++
    }

    return nums
}