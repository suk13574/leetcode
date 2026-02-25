
import "sort"

func sortByBits(arr []int) []int {
	sort.Ints(arr)

	countOneBit := func(n int) int {
		cnt := 0
		for n > 0 {
			cnt++
			n &= n - 1
		}
		return cnt
	}

	bits := make([][]int, 32)
	for i := 0; i < 32; i++ {
		bits[i] = []int{}
	}
	for _, n := range arr {
		cnt := countOneBit(n)

		bits[cnt] = append(bits[cnt], n)
	}

	res := []int{}
	for _, b := range bits {
		res = append(res, b...)
	}

	return res
}