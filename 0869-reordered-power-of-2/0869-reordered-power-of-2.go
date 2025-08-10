import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func reorderedPowerOf2(n int) bool {
	powerOf2 := make([]string, 32)
	for i := 0; i < 32; i++ {
		v := 1 << i
		str := strconv.Itoa(v)

		strs := strings.Split(str, "")
		sort.Strings(strs)
		sortedStr := strings.Join(strs, "")

		powerOf2[i] = sortedStr
	}

	str := strconv.Itoa(n)
	strs := strings.Split(str, "")
	sort.Strings(strs)
	sortedN := strings.Join(strs, "")

	return slices.Contains(powerOf2, sortedN)
}