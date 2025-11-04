import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	count := make(map[int]int)

	res := make([]int, 0, n-k+1)
	l := 0
	for r := 0; r < len(nums); r++ {
		count[nums[r]]++

		if r-l+1 == k {
			res = append(res, xSum(count, x))

			count[nums[l]]--
			l++
		}
	}
	return res
}

func xSum(count map[int]int, x int) int {
	type num struct {
		n   int
		cnt int
	}

	sortedNums := make([]num, len(count))
	for n, cnt := range count {
		sortedNums = append(sortedNums, num{n: n, cnt: cnt})
	}
	sort.Slice(sortedNums, func(i, j int) bool {
		if sortedNums[i].cnt == sortedNums[j].cnt {
			return sortedNums[i].n > sortedNums[j].n
		}
		return sortedNums[i].cnt > sortedNums[j].cnt
	})

	limit := min(len(sortedNums), x)

	sum := 0
	for i := 0; i < limit; i++ {
		sum += sortedNums[i].cnt * sortedNums[i].n
	}
	return sum
}