import (
	"fmt"
	"strconv"
)

const MOD = 1_000_000_007

func pow2mod(exp int) int {
	result := 1
	base := 2
	e := exp

	for e > 0 {
		if e&1 == 1 {
			result = int((int64(result) * int64(base)) % MOD)
		}
		base = int((int64(base) * int64(base)) % MOD)
		e = e >> 1
	}

	return result
}

func productQueries(n int, queries [][]int) []int {
	result := []int{}
	binaryStr := strconv.FormatInt(int64(n), 2)

	power := []int{}
	pos := 0

	for i := len(binaryStr) - 1; i >= 0; i-- {
		if binaryStr[i] == '1' {
			power = append(power, pos)
		}
		pos++
	}

	memo := make([]int, len(power)+1)
	for i := 0; i < len(power); i++ {
		memo[i+1] = memo[i] + power[i]
	}

	for _, query := range queries {
		i, j := query[0], query[1]

		value := pow2mod(memo[j+1] - memo[i])
		value = value % MOD
		result = append(result, int(value))
	}

	return result
}