func maxNumOfSubstrings(s string) []string {
    first := make([]int, 26)
    last := make([]int, 26)

    for i := 0; i < 26; i++ {
        first[i] = -1
        last[i] = -1
    }

    for i, ch := range s {
        idx := ch - 'a'

        if first[idx] == -1 {
            first[idx] = i
        }

        last[idx] = i
    }

    intervals := make([][2]int, 0, 26)
    for i := 0; i < 26; i++ {
        if first[i] == -1 {
            continue
        }

        valid := true
        left, right := first[i], last[i]

        for pos := left+1; pos <= right; pos++ {
            idx := s[pos] - 'a'

            if first[idx] < left {
                valid = false
                break
            }

            right = max(right, last[idx])
        }

        if valid {
            intervals = append(intervals, [2]int{left, right})
        }
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    res := []string{}
    lastEnd := -1

    for _, interval := range intervals {
        left, right := interval[0], interval[1]

        if left <= lastEnd {
            continue
        }

        res = append(res, s[left:right+1])
        lastEnd = right
    }

    return res
}