import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	noZeroInt := func(num int) bool {
		    return num > 0 && !strings.Contains(strconv.Itoa(num), "0")

	}

	result := make([]int, 0, 2)
	for i := 1; i < n; i++ {
		num1 := i
		num2 := n - i

		if noZeroInt(num1) && noZeroInt(num2) {
			result = append(result, num1)
			result = append(result, num2)
			break
		}
	}
	return result
}