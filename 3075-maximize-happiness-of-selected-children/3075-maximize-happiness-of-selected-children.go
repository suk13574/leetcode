import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Slice(happiness, func(i, j int) bool { return happiness[i] > happiness[j] })

    var res int64
    for i := 0; i < k; i++ {
        h := int64(happiness[i]) - int64(i)
        if h < 0 {
            break
        }
        res += h
    }

    return res
}
