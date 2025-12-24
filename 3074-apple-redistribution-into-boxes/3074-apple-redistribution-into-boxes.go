import "sort"

func minimumBoxes(apple []int, capacity []int) int {
    total := 0
    for _, a := range apple {
        total += a
    }   

    sort.Slice(capacity, func(i, j int) bool { return capacity[i] > capacity[j] })
    
    res := 0
    for _, c := range capacity {
        res++
        total -= c
        
        if total <= 0 {
            break
        }
    }

    return res
}