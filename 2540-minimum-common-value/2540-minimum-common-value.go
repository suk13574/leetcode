func getCommon(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	for idx1, idx2 := 0, 0; idx1 < n1 && idx2 < n2; {
		if nums1[idx1] == nums2[idx2] {
			return nums1[idx1]
		}

		if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			idx1++
		}
	}

	return -1
}