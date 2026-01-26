func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

	res := [][]int{}
    minDiff := 1 << 30
    for i := 1; i < len(arr); i++ {
        a := arr[i-1]
        b := arr[i]
        diff := b - a

        if diff < 0 {
            diff = -diff
        }

        if diff > minDiff {
            continue
        }
        if diff < minDiff {
            res = [][]int{}
            minDiff = diff
        }

        res = append(res, []int{a, b})
    }

    return res
}