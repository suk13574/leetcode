import (
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	kk := int64(k)
	count := 0
	last := int64(math.MinInt64)
	for i := 0; i < len(nums); i++ {
		num := int64(nums[i])
		lo := num - kk
		hi := num + kk

		now := lo
		if last+1 > now {
			now = last + 1
		}

		if now <= hi {
			last = now
			count++
		}
	}

	return count
}