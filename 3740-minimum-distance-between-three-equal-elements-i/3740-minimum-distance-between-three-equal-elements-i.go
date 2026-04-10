func minimumDistance(nums []int) int {
    n := len(nums)

    if n < 3 {
        return -1
    }

    tuples := make([][]int, 101)
    for i := 0; i < len(tuples); i++ {
        tuples[i] = []int{}
    }

    for idx, num := range nums {
        tuples[num] = append(tuples[num], idx)
    }

    res := 1 << 31
    for _, t := range tuples {
        if len(t) < 3 {
            continue
        }

        for i := 0; i+2 < len(t); i++ {
            d := (t[i+2]-t[i])*2
            res = min(res, d)
        }
    }

    if res == 1 << 31 {
        return -1
    }

    return res
}