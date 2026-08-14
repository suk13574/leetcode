func maxDistance(nums1 []int, nums2 []int) int {
    res := 0
    i := 0
    
    for j := 0; j < len(nums2); j++ {
        for i < len(nums1) && i <= j && nums1[i] > nums2[j] {
            i++
        }
        
        if i == len(nums1) {
            break
        }

        if nums1[i] <= nums2[j] {
            res = max(res, j-i)
        }
    }

    return res
}