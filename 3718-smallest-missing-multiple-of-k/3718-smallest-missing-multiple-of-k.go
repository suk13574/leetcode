func missingMultiple(nums []int, k int) int {
    sort.Ints(nums)

    need := k

    for _, num := range nums {
        if need > num {
            continue
        } else if need == num {
            need += k
        } else {
            return need  
        }
    }   
    
    return need
}