type NumArray struct {
    arr       []int
    prefixSum []int
}

func Constructor(nums []int) NumArray {
    p := make([]int, len(nums))
    
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        p[i] += sum
    }

    return NumArray{arr: nums, prefixSum: p}
}


func (na *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return na.prefixSum[right]
    }
    return na.prefixSum[right] - na.prefixSum[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */