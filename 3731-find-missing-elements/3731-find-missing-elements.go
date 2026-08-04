import "sort"

func findMissingElements(nums []int) []int {
	sort.Ints(nums)

	res := []int{}
	var need int
	for i, num := range nums {
		if i == 0 {
			need = num + 1
			continue
		}

		for need < num {
			res = append(res, need)
			need++
		}
		need++
	}

	return res
}