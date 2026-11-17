type Group struct {
	value int
	count int
}

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)

	dq := []Group{}

	var cost int64
	var res int64

	r := n - 1
	for l := n - 1; l >= 0; l-- {
		v := nums[l]
		count := 1

		for len(dq) > 0 && dq[len(dq)-1].value < v {
			last := dq[len(dq)-1]

			cost += int64(v-last.value) * int64(last.count)
			count += last.count

			dq = dq[:len(dq)-1]
		}

		dq = append(dq, Group{
			value: v,
			count: count,
		})

		for cost > int64(k) {
			cost -= int64(dq[0].value - nums[r])

			dq[0].count--

			if dq[0].count == 0 {
				dq = dq[1:]
			}

			r--
		}

		res += int64(r - l + 1)
	}

	return res
}