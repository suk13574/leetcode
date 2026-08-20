func resultArray(nums []int) []int {
	arr1 := make([]int, 1, len(nums)/2)
	arr2 := make([]int, 1, len(nums)/2)

	arr1[0] = nums[0]
	arr2[0] = nums[1]

	for i := 2; i < len(nums); i++ {
		n1 := len(arr1)
		n2 := len(arr2)

		if arr1[n1-1] > arr2[n2-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}

	return append(arr1, arr2...)
}