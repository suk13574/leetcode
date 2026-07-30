import "sort"

func minimumPushes(word string) int {
	n := len(word)

	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[word[i]-'a']++
	}

	sort.Ints(count)
	sort.Sort(sort.Reverse(sort.IntSlice(count)))

	res := 0
	availableKey := 8
	for i, c := range count {
		push := (i / availableKey) + 1
		res += c * push
	}

	return res
}