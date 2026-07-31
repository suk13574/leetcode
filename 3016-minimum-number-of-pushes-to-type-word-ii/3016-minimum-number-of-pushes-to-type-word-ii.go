func minimumPushes(word string) int {
    n := len(word)
    
    count := make([]int, 26)
    for i := 0; i < n; i++ {
        count[word[i]-'a']++
    }

    sort.Sort(sort.Reverse(sort.IntSlice(count)))

    res := 0
    availableKey := 8
    for i, c := range count {
        if c == 0 {
            break
        }

        push := (i / availableKey) + 1
        res += c * push
    }

    return res
}